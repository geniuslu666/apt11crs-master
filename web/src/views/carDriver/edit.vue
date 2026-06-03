<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="formValue.id > 0 ? driverName + '-' + '司机设置 #' + formValue.id : '添加司机'"
      >
        <n-tabs type="line" animated v-model:value="tabValue" v-if="formValue.id > 0">
          <n-tab-pane name="base" tab="基础信息">
            <EditBase ref="editBaseRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="settlement" tab="结算设置">
            <EditSettlement ref="editSettlementRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="qualification" tab="资质信息">
            <EditQualification ref="editQualificationRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
<!--          <n-tab-pane name="order" tab="预约设置">-->
<!--            <EditOrder ref="editOrderRef" :formData="formValue" @reloadInfo="load"/>-->
<!--          </n-tab-pane>-->
        </n-tabs>
<!--        <EditBase ref="editBaseRef" :formData="formValue" v-else/>-->
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          v-else
        >
          <n-tabs type="line" animated v-model:value="newTabValue">
            <n-tab-pane name="base" tab="基础信息">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="司机昵称" path="nickname">
                    <n-input placeholder="请输入司机昵称" v-model:value="formValue.nickname" style="width: 150px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="真实姓名" path="name">
                    <n-input placeholder="请输入真实姓名" v-model:value="formValue.name" style="width: 150px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="性别" path="sex" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.sex">
                      <n-radio-button
                        :value="1"
                        label="男"
                      />
                      <n-radio-button
                        :value="2"
                        label="女"
                      />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="手机号码" path="phone" :show-require-mark="true">
                    <n-input-group>
<!--                      <n-input-group-label>+</n-input-group-label>-->
                      <n-input placeholder="区号" v-model:value="formValue.phoneArea" style="width: 50px" />
                      <n-input-group-label>-</n-input-group-label>
                      <n-input placeholder="手机号" v-model:value="formValue.phone" style="width: 150px" />
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="照片" path="photo" :show-require-mark="true" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.photo" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：128px*174px</div>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="年龄" path="age" :show-require-mark="true">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.age" style="width: 100px" />
                      <n-input-group-label>岁</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="从业年数" path="workYears" :show-require-mark="true">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.workYears" style="width: 100px" />
                      <n-input-group-label>年</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="语言能力" path="language" :show-require-mark="true">
                    <n-checkbox-group v-model:value="formValue.languageArr">
                      <n-space>
                        <n-checkbox
                          v-for="item in languageArr"
                          :value="item.value"
                          :label="item.label"
                        />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="合作类型" path="cooperateTypeId" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.cooperateTypeId">
                      <n-radio-button
                        v-for="item in cooperateTypeList"
                        :key="item.id"
                        :value="item.id"
                        :label="item.typeName"
                      />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>
            <n-tab-pane name="settlement" tab="结算设置">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="结算模式" path="settlementId">
                    <n-select
                      v-model:value="formValue.settlementId"
                      :options="settlementList"
                      label-field="name"
                      value-field="id"
                      style="width: 200px"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="服务分成" path="settlementType">
                    <n-radio-group v-model:value="formValue.settlementType" name="settlementType">
                      <n-space>
                        <n-radio :value="1">
                          跟随配置
                        </n-radio>
                        <n-radio :value="2">
                          自定义
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item label="分成结算比例" path="orderConfirmHour" v-if="formValue.settlementType == 2">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :max="100" v-model:value="formValue.settlementRate" style="width: 100px" />
                      <n-input-group-label>%</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>
            <n-tab-pane name="qualification" tab="资质信息">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="资质材料" path="qualityMaterialsArr">
                    <FileChooser :maxNumber="10" v-model:value="formValue.qualityMaterialsArr" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>
          </n-tabs>
          <div style="text-align: center;margin-top: 30px">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted} from 'vue';
import {Edit, View} from '@/api/carDriver';
import {State, newState,loadOptions, options} from './model';
import {useRouter} from "vue-router";
import EditBase from "@/views/carDriver/edit_base.vue";
import EditQualification from "@/views/carDriver/edit_qualification.vue";
import EditSettlement from "@/views/carDriver/edit_settlement.vue"
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {List as SettlementList} from "@/api/carSettlement";
import {List as CooperateTypeList} from "@/api/carCooperateType";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";

const settingStore = useProjectSettingStore();
const newTabValue = ref('base')

const tabValue = ref('base')
const show = ref(false);
const router = useRouter();
const params = router.currentRoute.value.params;
const formValue = ref<State>(newState(null));
const editBaseRef = ref();
const editQualificationRef = ref();
const editSettlementRef = ref();
const driverName = ref('');
const languageArr = ref([
  {
    label: '中文',
    value: 'zh'
  },
  {
    label: '英文',
    value: 'en'
  },
  {
    label: '日语',
    value: 'ja'
  },
  {
    label: '韩语',
    value: 'ko'
  }
])

const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();
const settlementList = ref([]);
const cooperateTypeList = ref([]);
const tabsViewStore = useTabsViewStore();
const rules = ref({
  nickname: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入昵称'
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入真实姓名'
  },
});

async function loadSettlementList(){
  let settlementArrListOrg = await SettlementList({
    status: 1,
    Pagination: false
  })
  settlementList.value = settlementArrListOrg.list
}

async function loadCooperateTypeList(){
  let cooperateTypeArrListOrg = await CooperateTypeList({
    isThird: 2,
    status: 1,
    Pagination: false
  })
  cooperateTypeList.value = cooperateTypeArrListOrg.list
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.phoneArea || !formValue.value.phone){
        formBtnLoading.value = false;
        message.error('手机号请填写完整');
        return false;
      }
      if(!formValue.value.photo){
        formBtnLoading.value = false;
        message.error('请上传照片');
        return false;
      }
      if(!formValue.value.age){
        formBtnLoading.value = false;
        message.error('请输入年龄');
        return false;
      }
      if(!formValue.value.workYears){
        formBtnLoading.value = false;
        message.error('请输入从业年数');
        return false;
      }
      if(formValue.value.languageArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择语言能力');
        return false;
      }
      if(!formValue.value.cooperateTypeId){
        formBtnLoading.value = false;
        message.error('请选择合作类型');
        return false;
      }

      if(!formValue.value.settlementId){
        formBtnLoading.value = false;
        message.error('请选择结算模式');
        newTabValue.value = 'settlement'
        return false;
      }

      formValue.value.language = formValue.value.languageArr.join(',')
      formValue.value.qualityMaterials = formValue.value.qualityMaterialsArr.length > 0 ? formValue.value.qualityMaterialsArr.join(',') : ''

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          tabsViewStore.closeSignal('2');
          router.push({ name: 'carDriverIndex', params: {  } });
        }, 500);
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

function load() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    View({
      id: params.id,
      isLanguage: true
    })
      .then((res) => {
        formValue.value = res;
        driverName.value = res.nickname
      })
      .finally(() => {
        show.value = false;
      });
  });
}

onMounted(async() => {
  if(params.id > 0){
    load();
  }else{
    show.value = true;
    await loadOptions();
    await loadSettlementList()
    await loadCooperateTypeList()
    show.value = false;
  }
});
</script>

<style lang="less"></style>


