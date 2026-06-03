<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.status == 'OPENING' ? '开业原因' : '停业原因'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item :label="formValue.status == 'OPENING' ? '开业原因' : '停业原因'" path="desc">
                  <n-input
                    type="textarea"
                    :placeholder="formValue.status == 'OPENING' ? '请输入开业原因' : '请输入停业原因'"
                    v-model:value="formValue.desc"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import { Status } from '@/api/foodRestaurant';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { options } from './model';
import {adaModalWidth, getOptionLabel} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  status: '',
  desc: '',
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});
const rules = ref({
  desc: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: formValue.value.status == 'OPENING' ? '请输入开业原因' : '请输入停业原因',
  },
})

function openModal(id, status) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  formValue.value.status = status;
  loading.value = false;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Status({
        id: formValue.value.id,
        status: formValue.value.status,
        desc: formValue.value.desc
      }).then((_res) => {
        message.success('设为' + getOptionLabel(options.value.open_status, formValue.value.status) + '成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.id = 0;
  formValue.value.status = '';
  formValue.value.desc = '';
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


