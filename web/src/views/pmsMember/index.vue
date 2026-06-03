<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
        padding: '20px',
      }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员信息</text>
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
        <template #lastLoginSlot="{ model, field }">
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

      <BasicTable ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id"
        :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000" :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow" @update:sorter="handleUpdateSorter">
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space hidden md:inline-flex"
            v-if="hasPermission(['/pmsMember/export'])">
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


    <Edit ref="editRef" @reloadTable="reloadTable" />
    <sendemail ref="sendemailref" />
    <sendnotify ref="sendnotifyref" />
    <sendsms ref="sendsmsref" />
    <h5Link ref="h5Linkref" />

    <ChooseCoupon ref="chooseCouponRef" />

    <ChooseThCoupon ref="chooseThCouponRef" />

    <MemberCancel ref="memberCancelRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List, Status, Export } from '@/api/pmsMember';
import { columns, schemas, loadOptions, options } from './model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import Edit from './edit.vue';
import sendemail from '@/views/smjcomm/sendemail.vue';
import sendnotify from '@/views/pmsMember/sendnotify.vue';
import sendsms from '@/views/smjcomm/sendsms.vue';
import h5Link from '@/views/smjcomm/h5Link.vue';
import { useRouter } from "vue-router";
import ChooseCoupon from "@/views/pmsMember/chooseCouponType.vue";
import ChooseThCoupon from "@/views/pmsMember/chooseThCoupon.vue";
import MemberCancel from "@/views/pmsMember/member_cancel.vue";
import { useSorter } from "@/hooks/common";
import { ExportOutlined } from "@vicons/antd";
import { usePermission } from "@/hooks/web/usePermission";

const { hasPermission } = usePermission();
const router = useRouter();
const message = useMessage();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const checkedIds = ref([]);
const sendemailref = ref();
const sendnotifyref = ref();
const sendsmsref = ref();
const h5Linkref = ref();
const chooseCouponRef = ref();
const chooseThCouponRef = ref();
const memberCancelRef = ref();
const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

const actionColumn = reactive({
  width: 260,
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
          auth: ['/pmsMember/view'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/pmsMember/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/pmsMember/status'],
        },
        {
          label: '注销',
          onClick: handleCancel.bind(null, record),
          ifShow: () => {
            console.log('memberCancelArr', record.memberCancelArr);
            return record.memberCancelArr == null;
          },
          auth: ['/pmsWithdraw/disagreeStaff'],
          type: 'error'
        },
        // {
        //   label: '编辑',
        //   onClick: handleEdit.bind(null, record),
        //   auth: ['/pmsMember/edit'],
        // },
      ],
      dropDownActions: [
        {
          label: '发通知',
          key: 'mes',
        },
        {
          label: '发邮件',
          key: 'mail',
        },
        {
          label: '发短信',
          key: 'telmes',
        },
        {
          label: '发优惠券',
          key: 'coupon',
        },
        {
          label: '发礼品券',
          key: 'thcoupon',
        },
        {
          label: 'H5分销链接',
          key: 'h5_link',
        },
      ],
      select: (key) => {
        if (key === 'mes') {
          return tomessage(record);
        } else if (key === 'mail') {
          return goemail(record);
        } else if (key === 'telmes') {
          return gosend(record);
        } else if (key === 'coupon') {
          return sendCoupon(record);
        } else if (key === 'thcoupon') {
          return sendThCoupon(record);
        } else if (key === 'h5_link') {
          return goLink(record);
        }
      },
    });
  },
});
function gosend(aa) {
  if (aa.phone) {
    sendsmsref.value.openModal({
      phone: aa.phone,
      area_no: aa.phoneArea,
    });
  } else {
    message.error('未找到联络方式');
  }
}
function goemail(aa) {
  if (aa.mail) {
    sendemailref.value.openModal(aa.mail);
  } else {
    message.error('该会员无邮箱');
  }
}
function goLink(aa) {
  if (aa.id) {
    h5Linkref.value.openModal({
      id: aa.id,
    });
  } else {
    message.error('未找到联络方式');
  }
}
function sendCoupon(record) {
  chooseCouponRef.value.openModal(record.id);
}
function sendThCoupon(record) {
  chooseThCouponRef.value.openModal(record.id);
}
const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  if (sortStatesRef.value.length > 0) {
    var sort = "";
    sortStatesRef.value.forEach(element => {
      if (element.columnKey == 'exp') {
        sort = element.order
      }
    });
    return await List({ ...searchFormRef.value?.formModel, ...res, sort: sort });
  } else {
    return await List({ ...searchFormRef.value?.formModel, ...res });
  }
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}
function tomessage(aa) {
  sendnotifyref.value.openModal(aa.id);
}
// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
}

// 查看详情
function handleView(record: Recordable) {
  router.push({ name: 'pmsMember_view', params: { id: record.id } });
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

// 会员注销
function handleCancel(record: Recordable) {
  memberCancelRef.value.openModal(record);
}

// 导出
function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export(searchFormRef.value?.formModel);
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>
