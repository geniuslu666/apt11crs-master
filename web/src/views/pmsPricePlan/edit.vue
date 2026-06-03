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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑价格计划 #' + formValue.id : '新增价格计划' }}</div>
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
            require-mark-placement="right-hanging"
          >
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                价格计划设置
              </div>
              <n-grid :cols="1">
                <n-gi>
                    <n-form-item label="价格计划名称" path="planName">
                      <n-input placeholder="价格计划名称，仅后台查看" v-model:value="formValue.planName" style="width: 300px"/>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="适用取消政策" path="isCancel" :show-require-mark="true">
                      <n-radio-group v-model:value="formValue.isCancel" name="isCancel">
                        <n-space>
                          <n-radio value="N">
                            不可取消
                          </n-radio>
                          <n-radio value="Y">
                            灵活取消(依据平台配置的取消政策)
                          </n-radio>
                        </n-space>
                      </n-radio-group>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="该价格计划的限制预订条件" path="bookingDaysType">
                      <n-radio-group v-model:value="formValue.bookingDaysType" name="bookingDaysType">
                        <n-space>
                          <n-radio value="N">
                            否(跟随物业配置)
                          </n-radio>
                          <n-radio value="Y">
                            是
                          </n-radio>
                        </n-space>
                      </n-radio-group>
                    </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item path="bookingDays" v-if="formValue.bookingDaysType == 'Y'">
                    <template #label>
                      <span class="n-form-item-label__text">满足预订时间限制</span><span class="n-form-item-label__asterisk">&nbsp;*</span><span class="n-form-item-label__text">（单位：天）</span>
                    </template>
                    <n-input-group>
                      <n-input-number placeholder="最小天数" :min="1" :precision="0" v-model:value="formValue.bookingDays" style="width: 150px" />
                    </n-input-group>
                    <template #feedback v-if="formValue.bookingDays && propertyMinDaysNotice > formValue.bookingDays">
                      <div style="font-size: 12px; color: red">限制天数不能小于物业可预订最小天数</div>
                    </template>
                  </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="价格计划管理" path="priceMode" :show-feedback="false" :show-require-mark="true">
                      <n-radio-group v-model:value="formValue.priceMode" name="priceMode">
                        <n-space>
                          <n-radio value="-">
                            比Standard Rate便宜
                          </n-radio>
                          <n-radio value="+">
                            比Standard Rate贵
                          </n-radio>
                        </n-space>
                      </n-radio-group>
                    </n-form-item>
                    <n-form-item path="bookingDays" style="display: block;margin-top: 10px">
                      <n-input-group>
                        <n-input-group-label>{{ formValue.priceMode }}</n-input-group-label>
                        <n-input-number :min="0" :precision="0" v-model:value="formValue.planValue" style="width: 120px" />
                        <n-select v-model:value="formValue.priceStandard" :options="priceStandardOptions" style="width: 80px"/>
                      </n-input-group>
                      <template #feedback>对于这个价格计划，您想要设置比Standard Rate更便宜还是更贵?</template>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="关联房型管理" path="roomTypeId" style="margin-top: 24px">
                      <n-radio-group v-model:value="formValue.roomTypeId" name="roomTypeId">
                        <n-space>
                          <n-radio v-for="item in roomTypeArr"
                                   :key="item.id"
                                   :value="item.uid"
                                   :label="item.name">
                          </n-radio>
                        </n-space>
                      </n-radio-group>
                    </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                对外名称
              </div>
              <n-grid :cols="1">
                <n-gi>
                    <n-form-item label="对外名称_简体中文" path="planShowName_zh" :show-require-mark="true">
                      <n-input placeholder="简体中文的对外名称" v-model:value="nameLanguage.zh" style="width: 400px"/>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="对外名称_日本语" path="planShowName_ja" :show-require-mark="true">
                      <n-input placeholder="日本语的对外名称" v-model:value="nameLanguage.ja" style="width: 400px"/>
                    </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="对外名称_English" path="planShowName_en" :show-require-mark="true">
                    <n-input placeholder="英文的对外名称" v-model:value="nameLanguage.en" style="width: 400px"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="对外名称-한국어" path="planShowName_ko" :show-require-mark="true">
                      <n-input placeholder="韩文的对外名称" v-model:value="nameLanguage.ko" style="width: 400px"/>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="对外名称_繁体中文" path="planShowName_zh_CN" :show-require-mark="true">
                      <n-input placeholder="繁体中文的对外名称" v-model:value="nameLanguage.zh_CN" style="width: 400px"/>
                    </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                可见性
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="用户组可见性" path="memberGroupIdType" :show-feedback="false" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.memberGroupIdType" name="memberGroupIdType">
                      <n-space>
                        <n-radio value="ALL">
                          全部用户
                        </n-radio>
                        <n-radio value="PART">
                          部分用户组
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item path="memberGroupId" style="display: block;margin-top: 10px" v-if="formValue.memberGroupIdType == 'PART'" :show-feedback="false">
                    <n-checkbox-group v-model:value="formValue.memberGroupIdArr">
                      <n-space>
                        <n-checkbox
                          v-for="item in memberGroupArr"
                          :key="item.id"
                          :value="item.id"
                          :label="item.memberGroup"
                        />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
                <n-gi style="margin-top: 24px">
                  <n-form-item label="会员等级可见性" path="memberLevelIdType" :show-feedback="false" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.memberLevelIdType" name="memberLevelIdType">
                      <n-space>
                        <n-radio value="ALL">
                          全部等级
                        </n-radio>
                        <n-radio value="PART">
                          部分等级
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item path="memberLevelId" style="display: block;margin-top: 10px" v-if="formValue.memberLevelIdType == 'PART'">
                    <n-checkbox-group v-model:value="formValue.memberLevelIdArr">
                      <n-space>
                        <n-checkbox
                          v-for="item in memberLevelArr"
                          :key="item.id"
                          :value="item.id"
                          :label="item.levelName"
                        />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, reactive, computed} from 'vue';
import { Edit, View } from '@/api/pmsPricePlan';
import { State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";
import {List as RoomTypeList} from "@/api/pmsRoomType";
import {GroupAll} from "@/api/pmsMemberGroup";
import {All} from "@/api/pmsMemberLevel";
import {jsontoobj} from "@/utils/smjcomm";

const emit = defineEmits(['reloadTable']);
const loading = ref(false);
const showModal = ref(false);
const message = useMessage();
const settingStore = useProjectSettingStore();
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const roomTypeArr = ref([])
const memberGroupArr = ref([])
const memberLevelArr = ref([])
const propertyMinDaysNotice = ref(0);
const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});
const priceStandardOptions = ref([
  {
    label: '%',
    value: 'PERCENT'
  },
  {
    label: 'JPY',
    value: 'AMOUNT'
  }
])
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

async function openModal(state: State) {
  showModal.value = true;

  loading.value = true;
  await loadRoomTypeList(state);
  await loadMemberGroupList();
  await loadMemberLevelList();

  // 新增
  if (!state.id || state.id < 1) {
    formValue.value = newState(state);
    loading.value = false;

    // 对外名称
    nameLanguage.zh = '';
    nameLanguage.en = '';
    nameLanguage.ja = '';
    nameLanguage.ko = '';
    nameLanguage.zh_CN = '';

    return;
  }

  // 编辑
  await loadInfo(state)
  loading.value = false;
}

async function loadInfo(state){
  const dataInfo = await View({
    id: state.id
  })
  formValue.value = dataInfo;
  if(dataInfo.bookingDays > 0){
    formValue.value.bookingDaysType = "Y"
  }else{
    formValue.value.bookingDaysType = "N"
  }
  // 物业最小预定区间
  propertyMinDaysNotice.value = dataInfo.propertyDetail.minDaysNotice;
  if(dataInfo.planShowNameLanguage){
    dataInfo.planShowNameLanguage = jsontoobj(dataInfo.planShowNameLanguage);
    nameLanguage.zh = dataInfo.planShowNameLanguage.zh && dataInfo.planShowNameLanguage.zh.content ? dataInfo.planShowNameLanguage.zh.content : ''
    nameLanguage.en = dataInfo.planShowNameLanguage.en && dataInfo.planShowNameLanguage.en.content ? dataInfo.planShowNameLanguage.en.content : ''
    nameLanguage.ko = dataInfo.planShowNameLanguage.ko && dataInfo.planShowNameLanguage.ko.content ? dataInfo.planShowNameLanguage.ko.content : ''
    nameLanguage.ja = dataInfo.planShowNameLanguage.ja && dataInfo.planShowNameLanguage.ja.content ? dataInfo.planShowNameLanguage.ja.content : ''
    if(dataInfo.planShowNameLanguage.zh_CN){
      nameLanguage.zh_CN = dataInfo.planShowNameLanguage.zh_CN.content ? dataInfo.planShowNameLanguage.zh_CN.content : ''
    }else if(dataInfo.planShowNameLanguage.zh_cn){
      nameLanguage.zh_CN = dataInfo.planShowNameLanguage.zh_cn.content ? dataInfo.planShowNameLanguage.zh_cn.content : ''
    }else{
      nameLanguage.zh_CN = ''
    }
  }
  if(dataInfo.memberGroupId){
    formValue.value.memberGroupIdType = "PART"
  }else{
    formValue.value.memberGroupIdType = "ALL"
  }
  formValue.value.memberGroupIdArr = dataInfo.memberGroupId ? dataInfo.memberGroupId.split(',').map((item)=>{
    return parseInt(item)
  }) : [];
  if(dataInfo.memberLevelId){
    formValue.value.memberLevelIdType = "PART"
  }else{
    formValue.value.memberLevelIdType = "ALL"
  }
  formValue.value.memberLevelIdArr = dataInfo.memberLevelId ? dataInfo.memberLevelId.split(',').map((item)=>{
    return parseInt(item)
  }) : [];
}

async function loadRoomTypeList(state){
  const roomTypeArrList = await RoomTypeList({
    puid: state.propertyId,
    pagination: false
  })
  roomTypeArr.value = roomTypeArrList.list
}

async function loadMemberGroupList(){
  const memberGroupArrList = await GroupAll({})
  memberGroupArr.value = memberGroupArrList.list
}

async function loadMemberLevelList(){
  const memberLevelArrList = await All({})
  memberLevelArr.value = memberLevelArrList.list
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(formValue.value.bookingDaysType == 'Y' && (!formValue.value.bookingDays || formValue.value.bookingDays <= 0 || formValue.value.bookingDays < propertyMinDaysNotice.value)){
        formBtnLoading.value = false;
        if(formValue.value.bookingDays < propertyMinDaysNotice.value){
          message.error('限制最小天数不能小于物业可预订最小天数');
        }else{
          message.error('请填写满足预订时间限制最小天数，并且最小天数不能为0');
        }
        return false;
      }
      // if(!formValue.value.planValue){
      //   formBtnLoading.value = false;
      //   message.error('请完善价格计划管理');
      //   return false;
      // }
      formValue.value.planShowNameLanguage = nameLanguage
      if(formValue.value.memberGroupIdType == 'PART' && (formValue.value.memberGroupIdArr.length <= 0)){
        formBtnLoading.value = false;
        message.error('请选择用户组');
        return false;
      }
      if(formValue.value.memberLevelIdType == 'PART' && (formValue.value.memberLevelIdArr.length <= 0)){
        formBtnLoading.value = false;
        message.error('请选择用户等级');
        return false;
      }

      if(formValue.value.bookingDaysType == 'N'){
        // delete formValue.value.bookingDays
        formValue.value.bookingDays = null;
      }

      if(formValue.value.memberGroupIdType == 'ALL'){
        // delete formValue.value.memberGroupId
        formValue.value.memberGroupId = null;
      }else{
        formValue.value.memberGroupId = formValue.value.memberGroupIdArr.join(',')
      }

      if(formValue.value.memberLevelIdType == 'ALL'){
        // delete formValue.value.memberLevelId
        formValue.value.memberLevelId = null;
      }else{
        formValue.value.memberLevelId = formValue.value.memberLevelIdArr.join(',')
      }

      // delete formValue.value.bookingDaysType
      // delete formValue.value.memberGroupIdType
      // delete formValue.value.memberGroupIdArr
      // delete formValue.value.memberLevelIdType
      // delete formValue.value.memberLevelIdArr

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
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
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
.level-detail-div{
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      margin-bottom: 8px;
    }
    &-content{

    }
  }
}
</style>


