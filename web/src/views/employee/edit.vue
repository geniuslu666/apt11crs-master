<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑员工 #' + formValue.id : '添加员工' }}</div>
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
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="2" x-gap="17">
              <n-gi>
                <n-form-item label="员工姓名" path="name">
                  <n-input placeholder="请输入员工姓名" v-model:value="formValue.name" :style="{ width: '300px' }"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="所属部门" path="parentId">
                  <n-tree-select
                    v-model:value="formValue.departmentId"
                    :options="departmentTreeOptions"
                    key-field="id"
                    label-field="name"
                    children-field="children"
                    :default-expand-all="true"
                    placeholder="请选择"
                    clearable
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="手机号码" path="phone" :show-require-mark="true">
                  <n-input-group>
                    <n-input placeholder="区号" v-model:value="formValue.phoneArea" style="width: 50px" />
                    <n-input-group-label>-</n-input-group-label>
                    <n-input placeholder="手机号" v-model:value="formValue.phone" style="width: 150px" />
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="员工编号" path="email">
                  <n-input placeholder="请输入员工编号" v-model:value="formValue.employeeNo" :style="{ width: '300px' }"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="status in options.sys_normal_disable"
                      :key="status.value"
                      :value="status.value"
                      :label="status.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="备注" path="notice">
                  <n-input type="textarea" placeholder="备注" v-model:value="formValue.remark" style="width: 400px"/>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/employee';
import { options, State, newState, rules } from './model';
import Editor from '@/components/Editor/editor.vue';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import {GetDepartmentTree} from "@/api/employeeDepartment";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const departmentTreeOptions = ref([]);

async function openModal(state: State) {
  showModal.value = true;
  loading.value = true

  await loadDepartmentTree();

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);

    loading.value = false;
    return;
  }

  // 编辑
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

async function loadDepartmentTree() {
  try {
    const res = await GetDepartmentTree();
    departmentTreeOptions.value = res.list || [];
  } catch (error) {
    console.error('加载部门树失败:', error);
  }
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


