<template>
  <div>
    <div class="title-name">
      <!--      <div>{{ namegetlang(userStore.getuserPms.nameLanguage) }}</div>-->

      <div class="dropdown-container" style="width: 500px;">
        <!-- 自定义触发元素 -->
        <n-popover trigger="click" placement="bottom-start" :show-arrow="false" :overlap="false" raw>
          <template #trigger>
            <div class="trigger-div" @click="handleClick">
              {{ propertyName }}
              <!--              {{ namegetlang(userStore.getuserPms.nameLanguage) || '全部物业' }}-->
              <n-icon :component="isOpen ? ChevronUp : ChevronDown" />
            </div>
          </template>

          <!-- 下拉选择器 -->
          <n-select style="width: 350px;" v-model:value="selectedValue" filterable :filter="filterHandler"
            :options="sortedOptions" :render-label="renderLabel" value-field="id"
            :menu-props="{ style: { width: '350px' } }" @update:value="handleSelect" />
        </n-popover>
      </div>

    </div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
        padding: '20px',
      }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">系统入住订单</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
      padding: '0 20px 20px',
    }">
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable"
        @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #createdAtSlot="{ model, field }">
          <n-date-picker 
            v-model:formatted-value="model[field]" 
            type="datetimerange" 
            value-format="yyyy-MM-dd HH:mm:ss"
            clearable
            :shortcuts="defRangeShortcuts()"
            style="width: 100%"
          />
        </template>
      </BasicForm>
      <BasicTable ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id"
        :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000" :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space hidden md:inline-flex"
            v-if="hasPermission(['/pmsAppReservation/export'])">
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <!--    <Edit ref="editRef" @reloadTable="reloadTable" />-->
    <n-modal v-model:show="state.showreturn" :mask-closable="false" preset="dialog" title="退款" positive-text="提交"
      negative-text="算了" @positive-click="onPositiveClick" @negative-click="onNegativeClick" :style="{
        width: dialogWidth,
      }">
      <div style="width: 500px;max-width: 100%;" v-if="state.showreturn">
        <n-form ref="formRef" :model="state.returnAmount" :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="120" class="py-4">
          <n-form-item label="订单号" path="maxAmount">
            {{ state.returnAmount.ordersn }}
          </n-form-item>
          <n-form-item label="最大退款金额" path="maxAmount">
            {{ state.returnAmount.maxAmount }} (JPY)
          </n-form-item>
          <n-form-item label="主动退款" path="maxAmount">
            <n-input-number v-model:value="state.returnAmount.refundAmount" style="width: 220px"
              :max="state.returnAmount.maxAmount"
              :placeholder="'最大退款金额' + state.returnAmount.maxAmount"></n-input-number>
          </n-form-item>
          <n-form-item label="退款方式" path="isCancelOrder">
            <a-radio-group v-model:value="state.returnAmount.isCancelOrder" name="radioGroup">
              <a-radio value="N">仅退款</a-radio>
              <a-radio value="Y" v-if="state.orderStatus != 'CANCEL'">退款并取消订单</a-radio>
              <a-radio value="Y" v-else disabled="true">退款并取消订单</a-radio>
            </a-radio-group>
            <template #feedback>
              <div style="font-size: 12px">仅退款：订单会处于部分退款状态，订单状态不会改变</div>
              <div style="font-size: 12px; margin-bottom: 8px">退款并取消订单：退款后，订单会处于已取消状态</div>
            </template>
          </n-form-item>
          <n-form-item label="退款取消说明" path="reason">
            <n-input type="textarea" placeholder="请输入退款说明" v-model:value="state.returnAmount.reason" />
          </n-form-item>
        </n-form>
      </div>
    </n-modal>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import { computed, h, onMounted, reactive, ref, watch } from 'vue';
import { NButton, useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { Delete, Export, List, paycloud, Status } from '@/api/pmsAppReservation';
import { ExportOutlined } from '@vicons/antd';
import { columns, loadOptions, options, schemas } from './model';
import { adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import View from './view.vue';
import { useUserStore } from "@/store/modules/user";
import { useProjectSettingStore } from "@/store/modules/projectSetting";
import { ChevronDown, ChevronUp } from "@vicons/ionicons5";
import { storage } from "@/utils/Storage";
import { getlang } from "@/utils/smjcomm";

const userStore = useUserStore();
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const viewRef = ref();
const checkedIds = ref([]);
const settingStore = useProjectSettingStore();
const dialogWidth = computed(() => {
  return adaModalWidth(620);
});
const state = reactive({
  showreturn: false,
  btnloading: false,
  orderStatus: null,
  returnAmount: {
    ordersn: '',
    refundAmount: null,
    maxAmount: 0,
    isCancelOrder: "N",
    reason: ""
  }
})
const actionColumn = reactive({
  width: 165,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleEdit.bind(null, record),
          type: "primary"
        },
        {
          label: '退款',
          onClick: handleRefund.bind(null, record),
          type: "warning",
          ifShow: () => {
            if (record.isRefund === true) {
              if (record.refundAmount < record.orderAmount) {
                return true;
              }
            }
            return false;
          },
        },
      ],
      select: (key) => {
        if (key === 'view') {
          return handleView(record);
        }
      },
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 100,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.puid = userStore.getuserPms.uid ? userStore.getuserPms.uid : '';
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

//页面刷新  根据物业选择
watch(
  () => userStore.getuserPms,
  () => {
    reloadTable();  //当前页面重新加载的数据
  },
  {
    immediate: true,
    deep: true,
  }
);

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

const onPositiveClick = async () => {

  paycloud(state.returnAmount).then((_res) => {
    message.success('已提交');
    state.showreturn = false;
    state.returnAmount.refundAmount = null
    state.returnAmount.maxAmount = 0
    state.returnAmount.reason = ""
    state.returnAmount.isCancelOrder = "N"
    reloadTable();
  });

  /* await paycloud(state.returnAmount);
   message.success('已提交');
   state.showreturn = false;
   state.returnAmount.refundAmount = null
   // state.returnAmount.refundType = ""
   state.returnAmount.maxAmount = 0
   reloadTable();*/

};

function selectchange(key, options) {
  state.returnAmount.maxAmount = options.maxAmount;
  console.log(key, options);
}

const onNegativeClick = () => {
  state.returnAmount.refundAmount = null
  state.returnAmount.maxAmount = 0
  state.showreturn = false;
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleRefund(record: Recordable) {
  state.returnAmount.ordersn = record.orderSn;
  state.returnAmount.isCancelOrder = "N"
  state.returnAmount.maxAmount = 0
  state.orderStatus = record.orderStatus
  var payAmount = 0
  record.transactionDetail.map((m) => {
    if (m.payStatus == "DONE" && m.payType != "COUPON") {
      payAmount += m.payAmount;
    }
  });

  var refundAmount = 0
  if (record.transactionRefundDetail != null) {
    for (let i = 0; i < record.transactionRefundDetail.length; i++) {
      if (record.transactionRefundDetail[i].refundStatus == "DONE" || record.transactionRefundDetail[i].refundStatus == "WAIT") {
        refundAmount += record.transactionRefundDetail[i].refundAmount;
      }
    }
  }

  console.log("payAmount", payAmount)
  console.log("refundAmount", refundAmount)
  state.returnAmount.maxAmount += payAmount - refundAmount
  console.log("state.returnAmount", state.returnAmount)
  state.showreturn = true;
}

// 编辑数据
function handleEdit(record: Recordable) {
  viewRef.value.openModal(record);
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 单个删除
function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 批量删除
function handleBatchDelete() {
  if (checkedIds.value.length < 1) {
    message.error('请至少选择一项要删除的数据');
    return;
  }

  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: checkedIds.value }).then((_res) => {
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 导出
function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export(searchFormRef.value?.formModel);
}

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

const namegetlang = (data) => {
  if (data) {
    return getlang(data, userStore.language).content;
  } else {
    return '全部物业';
  }
};

// 物业名称
const propertyName = ref(userStore.getuserPms.id > 0 ? namegetlang(userStore.getuserPms.nameLanguage) : '全部物业');

// 物业列表
const wylists = ref(storage.get('wylists'));

// 控制箭头
const isOpen = ref(false)

// 下拉选中的值
const selectedValue = ref(userStore.getuserPms.id)

// 排序
const sortedOptions = computed(() => {

  // console.log("wylists.value", wylists.value)

  const selected = wylists.value.find(opt => opt.id === selectedValue.value)
  const others = wylists.value.filter(opt => opt.id !== selectedValue.value)
  if (selectedValue.value > 0) {
    return selected ? [wylists.value.find(opt => opt.id === 0), selected, ...others.filter(opt => opt.id !== 0)] : wylists.value
  }
  return wylists.value
})

// 自定义选项渲染
const renderLabel = (option) => {
  // console.log('option', namegetlang(option.nameLanguage))
  if (!option.id) return h('div', '全部物业')
  return h('div', { class: 'flex items-center gap-2' }, [
    h('span', namegetlang(option.nameLanguage))
  ])
}

// 筛选处理
const filterHandler = (inputValue, option) => {
  // 同时匹配label和value字段
  return namegetlang(option.nameLanguage).includes(inputValue) ||
    option.id.toString().includes(inputValue)
}

// 选择处理
const handleSelect = (value: string) => {
  // console.log('value', value)
  const selected = wylists.value.find(opt => opt.id === value)

  console.log('selected', selected)
  storage.set('userPms', selected);
  userStore.setuserPms(selected);
  propertyName.value = value > 0 ? namegetlang(selected.nameLanguage) : '全部物业'

  isOpen.value = false

}

// 点击触发元素
const handleClick = () => {
  // 可添加点击逻辑
  isOpen.value = !isOpen.value
}

onMounted(() => {
  wylists.value.unshift({ name: '全部', id: 0, uid: '' })
  loadOptions();
});
</script>

<style lang="less" scoped>
.title-name {
  margin: 15px 0 19px 20px;
  font-weight: 500;
  font-size: 30px;
  color: #3D3D3D;
  line-height: 42px;
}
</style>
