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
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">入住订单</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
      padding: '0 20px 20px',
    }">
      <div class="statusTab">
        <div :class="tabValue == 'CONFIRM' ? 'active' : ''"><span @click="handleUpdateValue('CONFIRM')">已确认</span></div>
        <div :class="tabValue == 'CANCEL' ? 'active' : ''"><span @click="handleUpdateValue('CANCEL')">已取消</span></div>
        <div :class="tabValue == 'ALL' ? 'active' : ''"><span @click="handleUpdateValue('ALL')">全部</span></div>
      </div>
      <BasicForm ref="searchFormRef" labelWidth="200" @register="register" @submit="reloadTable" @reset="reloadTable"
        @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable ref="actionRef" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn"
        :scroll-x="scrollX" :resizeHeightOffset="-10000">
        <template #tableTitle>

        </template>
      </BasicTable>
      <OrderView ref="orderViewRef" />
    </n-card>
    <Cancel ref="cancelRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted, watch } from 'vue';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { RoomList } from '@/api/pmsAppReservation';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import OrderView from "../pmsAppReservation/view.vue";
import { useUserStore } from "@/store/modules/user";
import Cancel from './cancel.vue';
import { ChevronDown, ChevronUp } from "@vicons/ionicons5";
import { storage } from "@/utils/Storage";
import { getlang } from "@/utils/smjcomm";

const userStore = useUserStore();
const actionRef = ref();
const searchFormRef = ref<any>({});
const orderViewRef = ref();
const cancelRef = ref();
const tabValue = ref('CONFIRM')

const actionColumn = reactive({
  width: 130,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          ifShow: () => {
            return record.orderSn;
          },
        },
        {
          label: '取消',
          onClick: handleCancel.bind(null, record),
          ifShow: () => {
            return record.status !== "cancelled";
          },
          auth: ['/pmsAppReservation/cancel'],
          type: 'error'
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:5 xl:5 2xl:5' },
  labelWidth: 120,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.puid = userStore.getuserPms.uid ? userStore.getuserPms.uid : '';
  res.selectStatus = tabValue.value;
  return await RoomList({ ...searchFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 拒绝提现申请
function handleCancel(record: Recordable) {
  cancelRef.value.openModal(record);
}

function handleUpdateValue(e) {
  tabValue.value = e;

  reloadTable()
}


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

// 查看详情
function handleView(record: Recordable) {
  record.viewtype = '住宿订单'
  orderViewRef.value.openModal(record);
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

  console.log("wylists.value", wylists.value)

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

.statusTab {
  display: flex;
  margin-bottom: 30px;

  div {
    padding: 0 16px;
    height: 28px;
    line-height: 28px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;

    span {
      cursor: pointer;
    }

    &.active {
      background: #F2F3F8;
      border-radius: 28px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
