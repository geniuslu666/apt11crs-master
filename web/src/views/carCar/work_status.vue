<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="设置工作状态"
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
                <n-form-item label="工作状态" path="workStatus">
                  <n-radio-group v-model:value="formValue.workStatus" name="workStatus">
                    <n-radio-button
                      v-for="type in workStatusOptions"
                      :key="type.value"
                      :value="type.value"
                      :label="type.label"
                    />
                  </n-radio-group>
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
import {ref} from 'vue';
import {View, WorkStatus} from '@/api/carCar';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { options } from './model';
import {getOptionLabel} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  workStatus: '',
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const rules = ref({
})

// const workStatusOptionsLabel = ref(options.value.driver_work_status);
const workStatusOptions = ref([
  {
    value: 'CANORDER',
    label: '空闲中'
  },
  {
    value: 'FAULT',
    label: '故障'
  }
]);

function openModal(id,status) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  formValue.value.workStatus = status;
  loading.value = false;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      WorkStatus({
        id: formValue.value.id,
        workStatus: formValue.value.workStatus,
      }).then((_res) => {
        message.success('设为' + getOptionLabel(options.value.car_work_status, formValue.value.workStatus) + '成功');
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
  formValue.value.workStatus = '';
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


