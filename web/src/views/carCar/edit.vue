<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :label-placement="settingStore.isMobile ? 'top' : 'left'"
        :label-width="180"
        class="py-4"
      >
        <n-card
          :bordered="false"
          class="proCard mt-4"
          size="small"
          :segmented="{ content: true }"
          :title="formValue.id > 0 ? '编辑服务 #' + formValue.id : '添加服务'"
        >
          <n-tabs type="line" animated v-model:value="tabValue">
            <n-tab-pane name="base" tab="基础信息">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="车辆名称" path="carName">
                    <n-input placeholder="请输入车辆名称" v-model:value="formValue.carName" style="width: 250px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型" path="typeId" :show-require-mark="true">
                    <n-select
                      v-model:value="formValue.typeId"
                      :options="carTypeList"
                      label-field="name"
                      value-field="id"
                      style="width: 200px"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车辆品牌型号" path="brand">
                    <n-input placeholder="请输入车辆品牌型号" v-model:value="formValue.brand" style="width: 250px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车牌号码" path="licenseNo">
                    <n-input placeholder="请输入车牌号码" v-model:value="formValue.licenseNo" style="width: 250px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="牌照颜色" path="licenseColor" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.licenseColor">
                      <n-radio-button
                        v-for="item in licenseColorList"
                        :key="item.value"
                        :value="item.value"
                        :label="item.label"
                      />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="排序" path="sort">
                    <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" style="width: 100px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="状态" path="status">
                    <n-radio-group v-model:value="formValue.status" name="status">
                      <n-radio-button
                        v-for="status in statusArr"
                        :key="status.value"
                        :value="status.value"
                        :label="status.label"
                      />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>

            <n-tab-pane name="content" tab="档案材料">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="车辆档案" path="qualityMaterialsArr" :show-require-mark="true">
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
        </n-card>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>

import {onMounted, ref} from "vue";
import {newState, State, rules} from "@/views/carCar/model";
import {useRouter} from "vue-router";
import {View,MaxSort,Edit} from "@/api/carCar";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Dicts} from "@/api/dict/dict";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {List as CarTypeList} from "@/api/carCarType";

const settingStore = useProjectSettingStore();
const tabValue = ref('base')
const show = ref(false);
const formBtnLoading = ref(false);
const router = useRouter();
const message = useMessage();
const tabsViewStore = useTabsViewStore();
const params = router.currentRoute.value.params;
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const carTypeList = ref([])
const statusArr = ref([]);
const licenseColorList = ref([
  {
    label: '白牌',
    value: 'WHITE'
  },
  {
    label: '绿牌',
    value: 'GREEN'
  },
  {
    label: '黄牌',
    value: 'YELLOW'
  }
])

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.typeId){
        formBtnLoading.value = false;
        message.error('请选择车型');
        return false;
      }
      formValue.value.qualityMaterials = formValue.value.qualityMaterialsArr.length > 0 ? formValue.value.qualityMaterialsArr.join(',') : ''
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'carCarIndex', params: {  } });
            }, 500);

          }
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

async function load() {
  const detailRes = await View({
    id: params.id,
  })
  formValue.value = detailRes
  formValue.value.qualityMaterialsArr = detailRes.qualityMaterials ? detailRes.qualityMaterials.split(',') : []
}

async function loadCarTypeList(){
  let carTypeArrListOrg = await CarTypeList({
    status: 1,
    Pagination: false
  })
  carTypeList.value = carTypeArrListOrg.list
}

async function loadDict(){
  let dictsList = await Dicts({
    types: ['sys_normal_disable'],
  })
  statusArr.value = dictsList.sys_normal_disable
}

async function loadMax(){
  let max = await MaxSort()
  formValue.value.sort = max.sort
}

async function init(){
  formValue.value = newState(null);
}

onMounted(async() => {
  show.value = true;
  await loadCarTypeList()
  await loadDict()
  if(parseInt(params.id) > 0){
    await load();
  }else{
    await init();
    await loadMax()
  }

  show.value = false;
})
</script>

<style lang="less"></style>


