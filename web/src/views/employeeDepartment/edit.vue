<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
          <template #header>
            <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ modalTitle }}</div>
          </template>
          <template #footer>
            <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
              取消
            </n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
              保存
            </n-button>
          </template>
          <n-spin :show="loading" description="请稍候...">
            <n-form
              :model="formParams"
              :rules="rules"
              ref="formRef"
              label-placement="top"
              label-width="auto"
            >
              <n-form-item label="上级部门" path="parentId">
                <n-tree-select
                  v-model:value="formParams.parentId"
                  :options="departmentTreeOptions"
                  key-field="id"
                  label-field="name"
                  children-field="children"
                  :default-expand-all="true"
                  placeholder="请选择上级部门，不选择则为顶级部门"
                  clearable
                  :disabled="isEditMode && formParams.id === formParams.parentId"
                />
              </n-form-item>
              <n-form-item label="部门名称" path="name">
                <n-input v-model:value="formParams.name" placeholder="请输入部门名称" />
              </n-form-item>
              <n-form-item label="状态" path="status">
                <n-radio-group v-model:value="formParams.status" name="status">
                  <n-radio-button
                    v-for="status in options.sys_normal_disable"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-form>
          </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useMessage } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";
import { State, newState, options } from './model';
import { Edit, View, GetDepartmentTree } from '@/api/employeeDepartment';

const emit = defineEmits(['reloadTable']);
const modalTitle = ref('');
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

const message = useMessage();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const formParams = ref<State>(newState(null));
const parentRecord = ref<any>(null);
const rules = ref({
  name: {
    required: true,
    message: '请输入部门名称',
  },
});
const departmentTreeOptions = ref([]);

const isEditMode = computed(() => {
  return formParams.value.id > 0;
});

async function openModal(record: State | null, title: string, parent?: any) {
  modalTitle.value = title;
  parentRecord.value = parent || null;
  resetForm();
  showModal.value = true;
  loading.value = true

  await loadDepartmentTree();

  if (record) {
    // 编辑模式
    const res = await View({ id: record.id });
    if (res) {
      formParams.value = newState(res);
    }
  } else if (parent) {
    // 新增子部门模式
    formParams.value.parentId = parent.id;
  }
  loading.value = false;
}

function resetForm() {
  formParams.value = newState(null);
  if (formRef.value?.restoreValidation) {
    formRef.value.restoreValidation();
  }
}

async function loadDepartmentTree() {
  try {
    const res = await GetDepartmentTree();
    departmentTreeOptions.value = res.list || [];
  } catch (error) {
    console.error('加载部门树失败:', error);
  }
}

function closeForm() {
  showModal.value = false;
}

function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formParams.value).then((_res) => {
          message.success(isEditMode.value ? '编辑成功' : '添加成功');
          setTimeout(() => {
            showModal.value = false;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  defineExpose({
    openModal,
  });
</script>
