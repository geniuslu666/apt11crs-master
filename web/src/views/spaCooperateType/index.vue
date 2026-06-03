<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">按摩营业类型</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
<!--          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/spaCooperateType/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List, Status } from '@/api/spaCooperateType';
import { columns, schemas, options, loadOptions } from './model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import Edit from './edit.vue';

const message = useMessage();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const checkedIds = ref([]);

const actionColumn = reactive({
  width: 288,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        // {
        //   label: '编辑',
        //   onClick: handleEdit.bind(null, record),
        //   auth: ['/spaCooperateType/edit'],
        // },
        //
        // {
        //   label: '禁用',
        //   onClick: handleStatus.bind(null, record, 2),
        //   ifShow: () => {
        //     return record.status === 1;
        //   },
        //   auth: ['/spaCooperateType/status'],
        // },
        // {
        //   label: '启用',
        //   onClick: handleStatus.bind(null, record, 1),
        //   ifShow: () => {
        //     return record.status === 2;
        //   },
        //   auth: ['/spaCooperateType/status'],
        // },
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
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  editRef.value.openModal(null);
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
}

// 单个删除
/*function handleDelete(record: Recordable) {
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
}*/

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

