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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑APP配置 #' + formValue.id : '添加APP配置' }}</div>
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
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="名称" path="name" >
                  <n-input placeholder="请输入名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="配置项" path="key">
                  <n-select placeholder="请选择配置项" v-model:value="formValue.key" :options="options.app_config_key" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="语言" path="language">
                  <n-select placeholder="请选择语言" v-model:value="formValue.language" :options="options.language" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="配置信息" path="value">
                  <Editor style="height: 300px" id="value" v-model:modelValue="formValue.value" />
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
  import { Edit, View } from '@/api/pmsAppconfig';
  import { options, State, newState } from './model';
  import Editor from '@/components/Editor/editor.vue';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import {rules} from "@/views/pmsChannel/model";

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

  function openModal(state: State) {
    showModal.value = true;
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      return;
    }
    // 编辑
    loading.value = true;
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
            closeForm();
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less">
.n-modal-container {
  z-index: 1000 !important;
}
</style>
