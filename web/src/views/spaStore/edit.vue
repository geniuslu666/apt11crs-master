<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :label-placement="settingStore.isMobile ? 'top' : 'left'"
        :label-width="150"
        class="py-4"
      >
        <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
          <n-gi span="1">
            <n-form-item label="门店名称" path="name">
              <n-input placeholder="请输入门店名称" v-model:value="formValue.name" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店负责人" path="headName">
              <n-input placeholder="请输入门店负责人" v-model:value="formValue.headName" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="区号" path="phoneArea">
              <n-input placeholder="请输入区号" v-model:value="formValue.phoneArea" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店电话" path="phone">
              <n-input placeholder="请输入门店电话" v-model:value="formValue.phone" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="营业时间" path="openTime">
              <n-input placeholder="请输入营业时间" v-model:value="formValue.openTime" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店地址" path="areaId" :show-require-mark="true">
              <n-tree-select
                v-model:value="formValue.areaId"
                :options="areaArr"
                key-field="id"
                label-field="areaName"
                :override-default-node-click-behavior="override"
                style="width: 250px"
                :show-path="true"
              />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店纬度" path="ggLat" :show-require-mark="true">
              <n-input placeholder="请输入门店纬度" v-model:value="formValue.ggLat" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="门店经度" path="ggLng" :show-require-mark="true">
              <n-input placeholder="请输入门店经度" v-model:value="formValue.ggLng" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label=" " path="detailAddress">
              <n-input placeholder="请输入详细地址" v-model:value="formValue.detailAddress" style="width: 300px"/>
            </n-form-item>
          </n-gi>
<!--          <n-gi span="1">-->
<!--            <div style="position: relative">-->
<!--              <div class="mt-3">-->
<!--                <googleMap-->
<!--                  :key="mapkey"-->
<!--                  :lat="formValue.ggLat"-->
<!--                  :long="formValue.ggLng"-->
<!--                  :heightmap="260"-->
<!--                  @addressData="addressDataGLoad($event)"-->
<!--                />-->
<!--              </div>-->
<!--            </div>-->
<!--          </n-gi>-->
        </n-grid>
      </n-form>
      <div style="text-align: center">
        <n-space justify="center">
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm">保存更新</n-button>
        </n-space>
      </div>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {Edit, Latest, View} from '@/api/spaStore';
import { State, newState, rules } from './model';
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {TreeSelectOverrideNodeClickBehavior, useMessage} from "naive-ui";
import {List as areaList} from "@/api/foodArea";
import googleMap from '@/views/smjcomm/GMap.vue';

const show = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const settingStore = useProjectSettingStore();
const message = useMessage();
const areaArr = ref([])
const override: TreeSelectOverrideNodeClickBehavior = ({ option }) => {
  if (option.children) {
    return 'toggleExpand'
  }
  return 'default'
}
let mapkey = ref(1);

const addressDataGLoad = (res) => {
  console.log('addressData', res);
  formValue.value.ggLat = res.lat + '';
  formValue.value.ggLng = res.lng + '';
};

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.areaId){
        formBtnLoading.value = false;
        message.error('请选择门店地址');
        return false;
      }
      if(!formValue.value.detailAddress){
        formBtnLoading.value = false;
        message.error('请输入详细地址');
        return false;
      }
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          load();
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

async function load() {
  const detailInfo = await Latest()
  if(detailInfo.id){
    formValue.value = detailInfo;
  }else{
    formValue.value = newState(null);
  }
}

async function loadAreaList(){
  const areaArrList = await areaList({
    status: 1,
    Pagination: false
  })
  areaArr.value = areaArrList.list
}

onMounted(async() => {
  show.value = true;
  await loadAreaList()
  await load();
  show.value = false;
});
</script>

<style lang="less">
.mapdiv {
  position: absolute;
  top: 15px;
  left: 15px;
  height: 38px;
  text-align: center;
  color: #3d3d3d;
  padding: 8px 15px;
  background-color: #fff;
  box-shadow: 4px 3px 5px #999999a8;
  z-index: 300;
  .tabview {
    color: #333;
    font-size: 14px;
    .showmap {
      color: #053dc8;
      font-weight: 550;
    }
    img {
      width: 18px;
      height: 18px;
      float: left;
      margin-right: 6px;
    }
  }
}
.mapbtm {
  position: absolute;
  z-index: 300;
  bottom: 0;
  text-align: left;
  width: 100%;
  background: #6a7b8c38;
  padding: 5px 10px;
  font-weight: 550;
}
</style>


