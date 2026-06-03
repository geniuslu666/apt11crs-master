<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="核账"
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
            :label-width="120"
            class="py-4"

          >
            <n-grid cols="2 s:1 m:1 l:2 xl:2 2xl:1" responsive="screen">
              <n-gi span="2" style="line-height: 18px">
                <n-form-item label="账期时间" path="startTime">
                  {{formValue.startTime}} ~ {{formValue.endTime}}
                </n-form-item>
              </n-gi>
              <n-gi span="2" v-if="formValue.settlementObject ==  'ISP'">
                <n-form-item label="服务商名称" path="ispName">
                  {{formValue.ispName}}
                </n-form-item>
              </n-gi>
              <n-gi span="2" v-if="formValue.settlementObject ==  'TECHNICIAN'">
                <n-form-item label="技师名称" path="technicianName">
                  {{formValue.technicianName}}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="订单金额" path="orderAmount">
                  {{formValue.orderAmount}}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="结算金额" path="settlementAmount">
                  {{formValue.settlementAmount}}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="核账人" path="settlementAmount">
                  {{formValue.operator}}
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="核账凭证" path="pic">
                  <UploadImage :maxNumber="1" v-model:value="formValue.verifyImg" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="核账说明" path="verifyDesc">
                  <n-input
                    type="textarea"
                    placeholder="请输入核账说明"
                    v-model:value="formValue.verifyDesc"
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
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { options } from './model';
import {adaModalWidth, getOptionLabel} from '@/utils/hotgo';
import {useUserStore} from "@/store/modules/user";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {Verify} from "@/api/spaSettlementOrder";

const emit = defineEmits(['reloadTable', 'reloadIspTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const userStore = useUserStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  settlementObject: "",
  ispId: 0,
  technicianId: 0,
  ispName: '',
  technicianName: '',
  orderAmount: 0,
  settlementAmount: 0,
  startTime: '',
  endTime: '',
  verifyImg: '',
  verifyDesc: '',
  operator: userStore.getUsername,
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
    message: '请输入核账说明',
  },
})

function openModal(record) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = record.id;
  formValue.value.settlementObject = record.settlementObject;
  formValue.value.ispId = record.ispId;
  formValue.value.technicianId = record.technicianId;
  if (record.settlementObject == "ISP"){
    formValue.value.ispName = record.ispDetail.name;
  }else{
    formValue.value.technicianName = record.technicianDetail.name;
  }
  formValue.value.orderAmount = record.orderAmount;
  formValue.value.settlementAmount = record.settlementAmount;
  formValue.value.startTime = record.startTime;
  formValue.value.endTime = record.endTime;
  loading.value = false;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Verify({
        id: formValue.value.id,
        verifyImg: formValue.value.verifyImg,
        verifyDesc: formValue.value.verifyDesc,
      }).then((_res) => {
        message.success('核账成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          if (formValue.value.settlementObject == "ISP"){
            emit('reloadIspTable');
          }else{
            emit('reloadTable');
          }
          // emit('reloadTable');
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
  formValue.value.verifyImg = '';
  formValue.value.verifyDesc = '';
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


