<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">司机管理</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm  ref="searchFormRef" @register="register" @submit="handleSubmit" @reset="handleReset" @keyup.enter="handleSubmit">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #createAtSlot="{ model, field }">
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
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane :name="item.value" :tab="item.label" v-for="item in options.sys_normal_disable">
        </n-tab-pane>
        <n-tab-pane :name="0" tab="全部">
        </n-tab-pane>
      </n-tabs>
      <BasicTable  ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/carDriver/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
<!--          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/carDriver/delete'])">-->
<!--            <template #icon>-->
<!--              <n-icon>-->
<!--                <DeleteOutlined />-->
<!--              </n-icon>-->
<!--            </template>-->
<!--            批量删除-->
<!--          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
    <WorkStatus ref="workStatusRef" @reloadTable="reloadTable"/>
    <Bind ref="bindRef" @reloadTable="reloadTable" />
    <BindCar ref="bindCarRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import {List,Status,Unbind} from '@/api/carDriver';
import { PlusOutlined } from '@vicons/antd';
import { columns, schemas, options, loadOptions } from './model';
import { adaTableScrollX,getOptionLabel } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import WorkStatus from '@/views/carDriver/work_status.vue';
import {useRouter} from "vue-router";
import Bind from "@/views/carDriver/bind.vue";
import BindCar from "@/views/carDriver/bind_car.vue";

const tabValue = ref(1)
const router = useRouter();
const message = useMessage();
const dialog = useDialog();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const checkedIds = ref([]);
const workStatusRef = ref();
const bindRef = ref();
const bindCarRef = ref();

const actionColumn = reactive({
  width: 400,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '司机设置',
          onClick: handleEdit.bind(null, record),
          auth: ['/carDriver/edit'],
        },
        {
          label: '工作状态',
          onClick: handleWorkStatus.bind(null, record),
          type: 'success',
          auth: ['/carDriver/workStatus'],
        },
        {
          label: '绑定用户',
          onClick: handleBind.bind(null, record),
          ifShow: () => {
            if(record.cooperateTypeDetail.isThird == 1){
              return false;
            }
            return record.memberId === 0;
          },
          auth: ['/carDriver/bind'],
        },
        {
          label: '解绑用户',
          onClick: handleUnbind.bind(null, record),
          ifShow: () => {
            if(record.cooperateTypeDetail.isThird == 1){
              return false;
            }
            return record.memberId > 0;
          },
          auth: ['/carDriver/unbind'],
          type: 'warning'
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/carDriver/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/carDriver/status'],
        },
        {
          label: record.carId > 0 ? '换绑车辆' : '绑定车辆',
          onClick: handleBindCar.bind(null, record),
          auth: ['/carDriver/bindCar'],
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  if(tabValue.value != 0){
    res.status = tabValue.value
  }
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

function handleSubmit(values: Recordable) {
  tabValue.value = values.status
  reloadTable()
}

function handleReset(values: Recordable) {
  tabValue.value = values.status
  reloadTable()
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e){
  tabValue.value = e;
  searchFormRef.value?.setFieldsValue({
    status: e
  });
  reloadTable()
}

// 添加数据
function addTable() {
  router.push({ name: 'carDriverEdit', params: { id: 0 } });
}

// 编辑数据
function handleEdit(record: Recordable) {
  router.push({ name: 'carDriverEdit', params: { id: record.id } });
}

// 修改状态
function handleWorkStatus(record: Recordable) {
  workStatusRef.value.openModal(record.id,record.workStatus);
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

// 绑定会员
function handleBind(record: Recordable) {
  bindRef.value.openModal(record.id);
}

// 绑定车辆
function handleBindCar(record: Recordable) {
  bindCarRef.value.openModal(record.id, record.carId);
}

// 解绑会员
function handleUnbind(record: Recordable) {
  let html = h(
    'div',
    null,
    [
      h(
        'div',
        null,
        {
          default: () => '当前绑定会员:',
        }
      ),
      h(
        'div',
        {
          style: {
            paddingLeft: '10px',
          }
        },
        {
          default: () => '会员号：' + record.memberDetail.memberNo,
        }
      ),
      h(
        'div',
        {
          style: {
            paddingLeft: '10px',
          }
        },
        {
          default: () => '会员姓名：' + record.memberDetail.fullName,
        }
      ),
      h(
        'div',
        {
          style: {
            paddingLeft: '10px',
          }
        },
        {
          default: () => '会员手机：' + record.memberDetail.phoneArea+"-"+record.memberDetail.phone,
        }
      ),
      h(
        'div',
        {
          style: {
            marginTop: '10px',
            color: 'red',
          }
        },
        {
          default: () => '你确定要解绑？解绑后，司机将无法使用司机APP登录，请谨慎操作！',
        }
      )
    ]
  )

  dialog.warning({
    title: '警告',
    content: () => {
      return html
    },
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Unbind(record).then((_res) => {
        message.success('解绑成功');
        reloadTable();
      });
    },
  });
}

// // 查看详情
// function handleView(record: Recordable) {
//   viewRef.value.openModal(record);
// }
//
// // 单个删除
// function handleDelete(record: Recordable) {
//   dialog.warning({
//     title: '警告',
//     content: '你确定要删除？',
//     positiveText: '确定',
//     negativeText: '取消',
//     onPositiveClick: () => {
//       Delete(record).then((_res) => {
//         message.success('删除成功');
//         reloadTable();
//       });
//     },
//   });
// }

// 批量删除
// function handleBatchDelete() {
//   if (checkedIds.value.length < 1){
//     message.error('请至少选择一项要删除的数据');
//     return;
//   }
//
//   dialog.warning({
//     title: '警告',
//     content: '你确定要批量删除？',
//     positiveText: '确定',
//     negativeText: '取消',
//     onPositiveClick: () => {
//       Delete({ id: checkedIds.value }).then((_res) => {
//         checkedIds.value = [];
//         message.success('删除成功');
//         reloadTable();
//       });
//     },
//   });
// }


onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

