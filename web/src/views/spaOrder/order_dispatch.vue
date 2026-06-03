<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">指派技师</div>
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
            <div class="dispatch-form-item-title">指定技师<span>*</span></div>
            <div class="dispatch-form-item-driver">
              <div class="dispatch-form-item-driver-info" v-if="!formValue.technicianIds">请选择技师</div>
              <block v-if="formValue.technicianIds">
                <div class="dispatch-form-item-driver-info choose" v-for="(item,index) in chooseTechnicianDetailInfo" :key="index">
                  <div>{{ item.name }}</div>
                  <div>{{ item.phone }}</div>
                </div>
              </block>
              <div class="dispatch-form-item-driver-btn">
                <n-button type="primary"  @click="chooseTechnicianBtn" v-if="!formValue.technicianIds">
                  <template #icon>
                    <n-icon>
                      <AddCircleOutline />
                    </n-icon>
                  </template>
                  选择技师
                </n-button>
                <n-button type="primary"  @click="chooseTechnicianBtn" v-else>
                  <template #icon>
                    <n-icon>
                      <SwapOutlined />
                    </n-icon>
                  </template>
                  更换技师
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
    <ChooseTechnician ref="chooseTechnicianRef" @reloadTechnician="handleChooseTechnicianInfo"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {Dispatch} from '@/api/spaOrder';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {useUserStore} from "@/store/modules/user";
import ChooseTechnician from "@/views/spaOrder/chooseTechnician.vue";
import {AddCircleOutline} from "@vicons/ionicons5";
import {SwapOutlined} from "@vicons/antd";

const emit = defineEmits(['reloadInfo']);
const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});

const message = useMessage();
const userStore = useUserStore();
const selectNum = ref(0);
const chooseTechnicianInfo = ref([]);
const chooseTechnicianDetailInfo = ref([]);
const chooseTechnicianName = ref([]);
const chooseTechnicianRef = ref();
const formValue = ref({
  id: 0,
  technicianIds: '',
  dispatchDesc: '',
  dispatchOperatorId: 0
});
const formBtnLoading = ref(false);

function openModal(id,num) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  selectNum.value = num;

  loading.value = false;
}

function chooseTechnicianBtn(){
  chooseTechnicianRef.value.openModal(formValue.value.id, selectNum.value, chooseTechnicianInfo.value.join(','));
}

function handleChooseTechnicianInfo(records){
  chooseTechnicianInfo.value = records.split(',')
  let chooseIds = []
  let chooseNames = []
  let chooseTechnicianDetailInfoOrg = []
  chooseTechnicianInfo.value.forEach((item)=>{
    let itemObj = item.split('-')
    let detailItem = {
      id: itemObj[0],
      name: itemObj[1],
      phone: itemObj[2] + ' ' + itemObj[3],
    }
    chooseTechnicianDetailInfoOrg.push(detailItem)
    chooseIds.push(itemObj[0]);
    chooseNames.push(itemObj[1]);
  })
  chooseTechnicianDetailInfo.value = chooseTechnicianDetailInfoOrg
  formValue.value.technicianIds = chooseIds.join(',')
  chooseTechnicianName.value = chooseNames
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.technicianIds){
    formBtnLoading.value = false;
    message.error('请选择技师');
    return false;
  }
  formValue.value.dispatchOperatorId = userStore.info ? userStore.info.id : 0;
  Dispatch(formValue.value).then((_res) => {
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
  selectNum.value = 0;
  formValue.value.id = 0;
  formValue.value.technicianIds = '';
  formValue.value.dispatchDesc = '';
  formValue.value.dispatchOperatorId = 0;
  chooseTechnicianDetailInfo.value = [];
  chooseTechnicianInfo.value = [];
  chooseTechnicianName.value = [];
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


