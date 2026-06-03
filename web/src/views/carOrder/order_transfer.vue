<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">指派司机</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;">
            取消
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;margin-left: 10px">
            确认
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="dispatch-form-item">
            <div class="dispatch-form-item-title">指定司机<span>*</span></div>
            <div class="dispatch-form-item-driver">
              <div class="dispatch-form-item-driver-info" v-if="parseInt(formValue.driverId) <= 0">请选择司机</div>
              <div class="dispatch-form-item-driver-info choose" v-else>
                <div>{{ chooseDriverInfo.nickname }}</div>
                <div>{{ chooseDriverInfo.phoneArea }} {{ chooseDriverInfo.phone }}</div>
              </div>
              <div class="dispatch-form-item-driver-btn">
                <n-button type="primary"  @click="chooseDriverBtn" v-if="parseInt(formValue.driverId) <= 0">
                  <template #icon>
                    <n-icon>
                      <AddCircleOutline />
                    </n-icon>
                  </template>
                  选择司机
                </n-button>
                <n-button type="primary"  @click="chooseDriverBtn" v-else>
                  <template #icon>
                    <n-icon>
                      <SwapOutlined />
                    </n-icon>
                  </template>
                  更换司机
                </n-button>
              </div>
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 35px">
            <div class="dispatch-form-item-title">调度说明</div>
            <div class="dispatch-form-item-textarea">
              <n-input
                type="textarea"
                placeholder="请输入调度说明"
                autosize style="min-height: 100px"
                v-model:value="formValue.dispatchDesc"
              />
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 20px">
            <div class="dispatch-form-item-title">操作人</div>
            <div class="dispatch-form-item-user">{{ userStore.username }}</div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <ChooseDriver ref="chooseDriverRef" @reloadDriver="handleChooseDriverInfo"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {Transfer} from '@/api/carOrder';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {useUserStore} from "@/store/modules/user";
import ChooseDriver from "@/views/carOrder/chooseDriver.vue";
import {AddCircleOutline} from "@vicons/ionicons5";
import {SwapOutlined} from  "@vicons/antd"

const emit = defineEmits(['reloadInfo']);
const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});

const message = useMessage();
const userStore = useUserStore();
const chooseDriverInfo = ref({});
const chooseDriverRef = ref();
const formValue = ref({
  id: 0,
  driverId: 0,
  dispatchDesc: '',
  dispatchOperatorId: 0
});
const formBtnLoading = ref(false);

function openModal(id) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  loading.value = false;
}


function chooseDriverBtn(){
  chooseDriverRef.value.openModal(formValue.value.id, formValue.value.driverId);
}

function handleChooseDriverInfo(record){
  chooseDriverInfo.value = record
  formValue.value.driverId = record.id
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.driverId){
    formBtnLoading.value = false;
    message.error('请选择司机');
    return false;
  }
  formValue.value.dispatchOperatorId = userStore.info ? userStore.info.id : 0;
  Transfer(formValue.value).then((_res) => {
    message.success('操作成功');
    formBtnLoading.value = false;
    setTimeout(() => {
      closeForm();
      emit('reloadInfo');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.id = 0;
  formValue.value.driverId = 0;
  formValue.value.dispatchDesc = '';
  formValue.value.dispatchOperatorId = 0;
}

defineExpose({
  openModal,
});
</script>

<style lang="less">
.dispatch-form-item{
  &-title{
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    span{
      color: #F73314;
    }
  }
  &-driver{
    margin-top: 8px;
    width: 405px;
    height: 80px;
    border: 1px solid #E0E0E6;
    border-radius: 2px;
    background: #FAFAFC;
    padding: 20px 14px;
    display: flex;
    justify-content: space-between;
    &-info{
      height: 100%;
      font-weight: 400;
      font-size: 14px;
      color: #C2C2C2;
      line-height: 20px;
      &.choose{
        font-weight: 500;
        display: flex;
        flex-direction: column;
        justify-content: center;
        color: #3D3D3D;
      }
    }
    &-btn{
      height: 100%;
      display: flex;
      align-items: center;
    }
  }
  &-textarea{
    margin-top: 8px;
    textarea{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
    }
  }
  &-user{
    margin-top: 8px;
    font-weight: 600;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
  }
}
</style>


