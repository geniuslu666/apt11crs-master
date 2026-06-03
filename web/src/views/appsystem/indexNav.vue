<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              首页导航配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="首页导航是否开启" path="isOpen">
                <n-switch v-model:value="formValue.isOpen" :unchecked-value="2" :checked-value="1"/>
                <template #feedback>关闭首页导航后，用户在APP端的首页不会显示出导航栏</template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isOpen == 1">
            <div style="margin-left: 200px;margin-top: 24px">
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
        <div style="text-align: center">
          <n-space justify="center">
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>

    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {computed, h, onMounted, reactive, ref} from 'vue';
import {NButton, useDialog, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';
import {BasicTable, TableAction} from "@/components/Table";
import {adaTableScrollX} from "@/utils/hotgo";
import { columns, loadOptions } from './nav_model';
import { List, Delete } from '@/api/pmsIndexNav';
import {PlusOutlined} from "@vicons/antd";
import Edit from "@/views/appsystem/navEdit.vue";

const actionRef = ref();
const dialog = useDialog();
const formBtnLoading = ref(false);
const group = ref('indexnav');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
const editRef = ref();

  const formValue = ref({
    isOpen: 1,
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

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        formBtnLoading.value = false;
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  onMounted(() => {
    loadOptions();
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
