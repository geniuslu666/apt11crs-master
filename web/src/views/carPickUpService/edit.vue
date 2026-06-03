<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="formValue.id > 0 ? '编辑接机服务 #' + formValue.id : '添加接机服务'"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="120"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="路线名称" path="serviceName">
                <n-input placeholder="请输入路线名称" v-model:value="formValue.serviceName" style="width: 250px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="车型" path="carTypeId" :show-require-mark="true">
                <n-select
                  v-model:value="formValue.carTypeId"
                  :options="carTypeList"
                  label-field="name"
                  value-field="id"
                  style="width: 200px"
                />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="出发机场" path="startIds" :show-require-mark="true">
                <n-select
                  v-model:value="formValue.startIds"
                  :options="startCarAddressList"
                  label-field="name"
                  value-field="id"
                  style="width: 200px"
                />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="目的地" path="endIdsType">
                <n-radio-group v-model:value="formValue.endIdsType" name="endIdsType">
                  <n-space>
                    <n-radio :value="1">
                      全部物业
                    </n-radio>
                    <n-radio :value="2">
                      部分物业
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item label=" " path="endIdsArr" v-if="formValue.endIdsType == 2">
                <n-select
                  placeholder="请选择目的地"
                  v-model:value="formValue.endIdsArr"
                  :options="endCarAddressList"
                  label-field="name"
                  value-field="id"
                  clearable
                  filterable
                  multiple
                  style="width: 400px"
                />
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预估距离" path="distance">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.distance" style="width: 100px" />
                  <n-input-group-label>KM</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预估时长" path="useTime">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.useTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="基础费用" path="price">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" v-model:value="formValue.price" style="width: 100px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="划线价" path="linePrice">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" v-model:value="formValue.linePrice" style="width: 100px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
                <template #feedback>不填或输入0则不显示划线价</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1" style="margin-top: 24px">
              <n-form-item label="免费等待时长" path="freeWaitTime">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.freeWaitTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="最大等待时间" path="maxWaitTime">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.maxWaitTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="超时计费" path="timeout">
                <n-input-group>
                  <n-input-group-label>每</n-input-group-label>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.timeoutPreTime" style="width: 100px" />
                  <n-input-group-label>分钟计费</n-input-group-label>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" v-model:value="formValue.timeoutPrePrice" style="width: 100px" />
                  <n-input-group-label>JPY，不足按上限计算</n-input-group-label>
                </n-input-group>
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
            <n-gi span="1">
              <n-form-item label="排序" path="sort">
                <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" style="width: 100px"/>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
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
import {onMounted, ref} from 'vue';
import { Edit, View, MaxSort } from '@/api/carService';
import { State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {List as CarTypeList} from "@/api/carCarType";
import {List as CarAddressList} from "@/api/carAddress";
import {useRouter} from "vue-router";
import {Dicts} from "@/api/dict/dict";
import {useTabsViewStore} from "@/store/modules/tabsView";


const emit = defineEmits(['reloadTable']);
const tabsViewStore = useTabsViewStore();
const message = useMessage();
const router = useRouter();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const params = router.currentRoute.value.params;
const carTypeList = ref([])
const startCarAddressList = ref([])
const endCarAddressIds = ref([])
const endCarAddressList = ref([])
const statusArr = ref([]);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.carTypeId){
        formBtnLoading.value = false;
        message.error('请选择车型');
        return false;
      }
      if(!formValue.value.startIds){
        formBtnLoading.value = false;
        message.error('请选择出发机场');
        return false;
      }
      if(formValue.value.endIdsType == 2 && formValue.value.endIdsArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择目的地');
        return false;
      }

      formValue.value.serviceType = 'PICKUP'
      if(formValue.value.endIdsType == 1){
        formValue.value.endIds = '';
      }else{
        formValue.value.endIds = formValue.value.endIdsArr.join(',')
      }
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          tabsViewStore.closeSignal('2');
          router.push({ name: 'carPickUpServiceIndex', params: {  } });
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

async function init(){
  formValue.value = newState(null);
  MaxSort()
    .then((res) => {
      formValue.value.sort = res.sort;
    })
}

async function loadCarTypeList(){
  let carTypeArrListOrg = await CarTypeList({
    status: 1,
    Pagination: false
  })
  carTypeList.value = carTypeArrListOrg.list
}

async function loadCarAddressList(){
  let carAddressArrListOrg = await CarAddressList({
    Pagination: false
  })

  if(carAddressArrListOrg.list.length > 0){
    carAddressArrListOrg.list.forEach((item) => {
      if(item.typeId == 1){
        if(item.status == 1){
          item.disabled = false;
        }else{
          item.disabled = true;
        }
        startCarAddressList.value.push(item)
      }
      if(item.typeId == 4){
        if(item.status == 1){
          endCarAddressIds.value.push(item.id)
          endCarAddressList.value.push(item)
        }
      }
    })
  }
}

async function load() {
  const detailRes = await View({
    id: params.id,
  })
  formValue.value = detailRes
  formValue.value.startIds = parseInt(detailRes.startIds)
  if(detailRes.endIds){
    let endIdsArr = detailRes.endIds.split(',').map((item) => {
      return parseInt(item)
    });
    formValue.value.endIdsArr = endIdsArr.filter(num => endCarAddressIds.value.includes(num));
    formValue.value.endIdsType = 2
  }else{
    formValue.value.endIdsArr = [];
    formValue.value.endIdsType = 1
  }
}

async function loadDict(){
  let dictsList = await Dicts({
    types: ['sys_normal_disable'],
  })
  statusArr.value = dictsList.sys_normal_disable
}

onMounted(async() => {
  loading.value = true;
  await loadCarTypeList()
  await loadCarAddressList()
  await loadDict()
  if(parseInt(params.id) > 0){
    await load();
  }else{
    await init();
  }

  loading.value = false;
})
</script>

<style lang="less"></style>


