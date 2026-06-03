<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <div style="margin-top: 24px">
              <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
                <template #tableTitle>
                  <n-button type="primary"  @click="addTable" class="min-left-space">
                    <template #icon>
                      <n-icon>
                        <PlusOutlined />
                      </n-icon>
                    </template>
                    添加
                  </n-button>
                </template>
              </BasicTable>
            </div>
          </n-gi>
        </n-grid>
      </n-form>
    </n-spin>

    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {computed, h, onMounted, reactive, ref} from 'vue';
import {NButton, useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from "@/components/Table";
import {adaTableScrollX} from "@/utils/hotgo";
import { columns, loadOptions } from './test_nav_model';
import { List, Delete } from '@/api/pmsTestNav';
import {PlusOutlined} from "@vicons/antd";
import Edit from "@/views/appsystem/testNavEdit.vue";

const actionRef = ref();
const dialog = useDialog();
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
const editRef = ref();

  const formValue = ref({
  });

const actionColumn = reactive({
  width: 288,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

// 加载表格数据
const loadDataTable = async (res) => {
  return await List({ ...res });
};

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

  const rules = {
  };

  onMounted(() => {
    loadOptions();
  });
</script>
