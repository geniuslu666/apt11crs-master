<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" @mask-click="closeForm">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">绑定用户</div>
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
            <div class="dispatch-form-item-title">是否是店长</div>
            <div class="dispatch-form-item-textarea">
              <n-radio-group v-model:value="formValue.isLeader" name="radioGroup">
                <n-radio :value="1">是</n-radio>
                <n-radio :value="2">否</n-radio>
              </n-radio-group>
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 35px">
            <div class="dispatch-form-item-title">绑定用户</div>
            <div class="dispatch-form-item-driver">
              <div class="dispatch-form-item-driver-info" v-if="parseInt(formValue.memberId) <= 0">请选择用户</div>
              <div class="dispatch-form-item-driver-info choose" v-else>
                <div>{{ chooseMemberInfo.fullName }}</div>
                <div>{{ chooseMemberInfo.memberNo }}</div>
                <div>{{ chooseMemberInfo.phoneArea }} {{ chooseMemberInfo.phone }}</div>
              </div>
              <div class="dispatch-form-item-driver-btn">
                <n-button type="primary"  @click="chooseMemberBtn" v-if="parseInt(formValue.memberId) <= 0">
                  <template #icon>
                    <n-icon>
                      <AddCircleOutline />
                    </n-icon>
                  </template>
                  选择用户
                </n-button>
                <n-button type="primary"  @click="chooseMemberBtn" v-else>
                  <template #icon>
                    <n-icon>
                      <SwapOutlined />
                    </n-icon>
                  </template>
                  更换用户
                </n-button>
              </div>
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>


    <ChooseMember ref="chooseMemberRef" @reloadMember="handleChooseMemberInfo"/>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import {Bind,} from '@/api/spaTechnician';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import {AddCircleOutline} from "@vicons/ionicons5";
import {SwapOutlined} from "@vicons/antd";
import ChooseMember from "@/views/spaTechnician/chooseMember.vue";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const chooseMemberInfo = ref({});
const chooseMemberRef = ref();
const formValue = ref({
  id: 0,
  isLeader: 2,
  memberId: 0
});
const formBtnLoading = ref(false);

function openModal(id) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  loading.value = false;
}

function chooseMemberBtn(){
  chooseMemberRef.value.openModal(formValue.value.id, formValue.value.memberId);
}

function handleChooseMemberInfo(record){
  chooseMemberInfo.value = record
  formValue.value.memberId = record.id
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.memberId){
    formBtnLoading.value = false;
    message.error('请选择用户');
    return false;
  }
  Bind(formValue.value).then((_res) => {
    message.success('操作成功');
    formBtnLoading.value = false;
    setTimeout(() => {
      closeForm();
      emit('reloadTable');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.id = 0;
  formValue.value.isLeader = 2;
  formValue.value.memberId = 0;
  chooseMemberInfo.value = {}
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


