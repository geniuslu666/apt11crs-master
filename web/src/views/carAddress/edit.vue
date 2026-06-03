<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="formValue.id > 0 ? '编辑地址 #' + formValue.id : '添加地址'"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="200"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="地点类型" path="typeId" :show-require-mark="true">
                <n-radio-group v-model:value="formValue.typeId">
                  <n-radio-button
                    v-for="item in typeArr"
                    :key="item.id"
                    :value="item.id"
                    :label="item.typeName"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="状态" path="status">
                <n-radio-group v-model:value="formValue.status" name="status">
                  <n-radio-button
                    v-for="status in options.sys_normal_disable"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="地点名称（后台）" path="name">
                <n-input placeholder="请输入地点名称（后台）" v-model:value="formValue.name" style="width: 200px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="机场航站楼" path="terminalName" v-if="formValue.typeId == 1">
                <n-input placeholder="请输入" v-model:value="formValue.terminalName" style="width: 100px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="机场代码" path="airportCode" v-if="formValue.typeId == 1">
                <n-input placeholder="请输入" v-model:value="formValue.airportCode" style="width: 100px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1" v-if="formValue.typeId == 4">
              <n-form-item label="选择物业" path="propertyIdsArr" :show-require-mark="true">
                <n-select
                  placeholder="请选择物业"
                  v-model:value="formValue.propertyId"
                  :options="propertyList"
                  label-field="name"
                  value-field="id"
                  clearable
                  filterable
                  style="width: 150px"
                  @update:value="handleUpdate"
                />
              </n-form-item>
            </n-gi>
            <template v-if="formValue.typeId != 4">
              <n-gi span="1">
                <n-form-item label="地点名称(app)_简体中文" path="subName_zh" :show-require-mark="true">
                  <n-input placeholder="请输入地点名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地点名称(app)_English" path="subName_en" :show-require-mark="true">
                  <n-input placeholder="请输入地点名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地点名称(app)_日本语" path="subName_ja" :show-require-mark="true">
                  <n-input placeholder="请输入地点名称" v-model:value="nameLanguage.ja" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="详细地址_简体中文" path="detailAddress_zh" :show-require-mark="true">
                  <n-input placeholder="请输入详细地址" v-model:value="addressLanguage.zh" @blur="translate('address','zh',addressLanguage.zh)" style="width: 40%"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="详细地址_English" path="detailAddress_en" :show-require-mark="true">
                  <n-input placeholder="请输入详细地址" v-model:value="addressLanguage.en" @blur="translate('address','en',addressLanguage.en)" style="width: 40%"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="详细地址_日本语" path="detailAddress_ja" :show-require-mark="true">
                  <n-input placeholder="请输入详细地址" v-model:value="addressLanguage.ja" style="width: 40%"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地点纬度" path="ggLat" :show-require-mark="true">
                  <n-input placeholder="请输入地点纬度" v-model:value="formValue.ggLat" style="width: 250px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地点经度" path="ggLng" :show-require-mark="true">
                  <n-input placeholder="请输入地点经度" v-model:value="formValue.ggLng" style="width: 250px"/>
                </n-form-item>
              </n-gi>
<!--              <n-gi span="1">-->
<!--                <div style="position: relative">-->
<!--                  <div class="mt-3">-->
<!--                    <googleMap-->
<!--                      :key="mapkey"-->
<!--                      :lat="formValue.ggLat"-->
<!--                      :long="formValue.ggLng"-->
<!--                      :heightmap="260"-->
<!--                      @addressData="addressDataGLoad($event)"-->
<!--                    />-->
<!--                  </div>-->
<!--                </div>-->
<!--              </n-gi>-->
            </template>
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
import {ref, computed, reactive, onMounted} from 'vue';
import { Edit, View } from '@/api/carAddress';
import {options, State, newState, rules, loadOptions} from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {Text} from "@/api/translate";
import googleMap from '@/views/smjcomm/GMap.vue';
import {List as typeList} from "@/api/carAddressType";
import {jsontoobj} from "@/utils/smjcomm";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {All as propertyAll, View as propertyView} from "@/api/pmsProperty";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const router = useRouter();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const params = router.currentRoute.value.params;
const propertyList = ref([]);

let mapkey = ref(1);
const typeArr = ref([])
const tabsViewStore = useTabsViewStore();

const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const addressLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

// 翻译
function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.ko = text
      return
    }
  }
  if(type == 'address'){
    if(lang == 'zh'){
      // 简体中文
      addressLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      addressLanguage.ko = text
      return
    }
  }
  if(text){
    if(lang == 'en'){
      text = text.toLowerCase()
    }
    Text({
      text: text
    }).then((_res) => {
      if(type == 'name'){
        if(lang == 'zh'){
          // 简体中文
          nameLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          // nameLanguage.ko = _res.ko
        }
      }
      if(type == 'address'){
        if(lang == 'zh'){
          // 简体中文
          addressLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          // addressLanguage.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

const addressDataGLoad = (res) => {
  console.log('addressData', res);
  formValue.value.ggLat = res.lat + '';
  formValue.value.ggLng = res.lng + '';
};

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  // 优惠券名称-多语言
  formValue.value.nameLanguage = nameLanguage
  // 优惠券使用说明-多语言
  formValue.value.addressLanguage = addressLanguage

  formValue.value.subName = ""
  formValue.value.detailAddress = ""

  console.log('formValue', formValue.value)
  // return

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          tabsViewStore.closeSignal('2');
          router.push({ name: 'carAddressIndex', params: {  } });
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

async function loadTypeList(){
  const typeArrList = await typeList({
    status: 1,
    Pagination: false
  })
  typeArr.value = typeArrList.list
}

async function loadPropertyList(){
  let propertyArrList = await propertyAll({
    isFindClose: true
  })
  propertyList.value = propertyArrList.list

}

function handleUpdate(value){
  console.log('propertyList', propertyList)
  console.log('value', value)
  // 查询物业详情
  loading.value = true;
  propertyView({ id: value })
    .then((res) => {
      // formValue.value = res;
      // let propertyDetail = res
      console.log("propertyDetail",  res)
      formValue.value.propertyId = res.id
      if(res.nameLanguage){
        res.nameLanguage = jsontoobj(res.nameLanguage);
        nameLanguage.zh = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
        nameLanguage.en = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
        nameLanguage.ko = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
        nameLanguage.ja = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
        if(res.nameLanguage.zh_CN){
          nameLanguage.zh_CN = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
        }else if(res.nameLanguage.zh_cn){
          nameLanguage.zh_CN = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
        }else{
          nameLanguage.zh_CN = ''
        }
      }

      if(res.addressLanguage){
        res.addressLanguage = jsontoobj(res.addressLanguage);
        addressLanguage.zh = res.addressLanguage.zh && res.addressLanguage.zh.content ? res.addressLanguage.zh.content : ''
        addressLanguage.en = res.addressLanguage.en && res.addressLanguage.en.content ? res.addressLanguage.en.content : ''
        addressLanguage.ko = res.addressLanguage.ko && res.addressLanguage.ko.content ? res.addressLanguage.ko.content : ''
        addressLanguage.ja = res.addressLanguage.ja && res.addressLanguage.ja.content ? res.addressLanguage.ja.content : ''
        if(res.addressLanguage.zh_CN){
          addressLanguage.zh_CN = res.addressLanguage.zh_CN.content ? res.addressLanguage.zh_CN.content : ''
        }else if(res.addressLanguage.zh_cn){
          addressLanguage.zh_CN = res.addressLanguage.zh_cn.content ? res.addressLanguage.zh_cn.content : ''
        }else{
          addressLanguage.zh_CN = ''
        }
      }

      formValue.value.ggLat = res.lat + '';
      formValue.value.ggLng = res.lng + '';
      formValue.value.lat = res.lat + '';
      formValue.value.lng = res.lng + '';

    })
    .finally(() => {
      loading.value = false;
    });
}

async function load() {
  // loading.value = true;
  const detailRes = await View({
    id: params.id,
    isLanguage: true
  })
  formValue.value = detailRes
  if(detailRes.nameLanguage){
    detailRes.nameLanguage = jsontoobj(detailRes.nameLanguage);
    nameLanguage.zh = detailRes.nameLanguage.zh && detailRes.nameLanguage.zh.content ? detailRes.nameLanguage.zh.content : ''
    nameLanguage.en = detailRes.nameLanguage.en && detailRes.nameLanguage.en.content ? detailRes.nameLanguage.en.content : ''
    nameLanguage.ko = detailRes.nameLanguage.ko && detailRes.nameLanguage.ko.content ? detailRes.nameLanguage.ko.content : ''
    nameLanguage.ja = detailRes.nameLanguage.ja && detailRes.nameLanguage.ja.content ? detailRes.nameLanguage.ja.content : ''
    if(detailRes.nameLanguage.zh_CN){
      nameLanguage.zh_CN = detailRes.nameLanguage.zh_CN.content ? detailRes.nameLanguage.zh_CN.content : ''
    }else if(detailRes.nameLanguage.zh_cn){
      nameLanguage.zh_CN = detailRes.nameLanguage.zh_cn.content ? detailRes.nameLanguage.zh_cn.content : ''
    }else{
      nameLanguage.zh_CN = ''
    }
  }

  if(detailRes.addressLanguage){
    detailRes.addressLanguage = jsontoobj(detailRes.addressLanguage);
    addressLanguage.zh = detailRes.addressLanguage.zh && detailRes.addressLanguage.zh.content ? detailRes.addressLanguage.zh.content : ''
    addressLanguage.en = detailRes.addressLanguage.en && detailRes.addressLanguage.en.content ? detailRes.addressLanguage.en.content : ''
    addressLanguage.ko = detailRes.addressLanguage.ko && detailRes.addressLanguage.ko.content ? detailRes.addressLanguage.ko.content : ''
    addressLanguage.ja = detailRes.addressLanguage.ja && detailRes.addressLanguage.ja.content ? detailRes.addressLanguage.ja.content : ''
    if(detailRes.addressLanguage.zh_CN){
      addressLanguage.zh_CN = detailRes.addressLanguage.zh_CN.content ? detailRes.addressLanguage.zh_CN.content : ''
    }else if(detailRes.addressLanguage.zh_cn){
      addressLanguage.zh_CN = detailRes.addressLanguage.zh_cn.content ? detailRes.addressLanguage.zh_cn.content : ''
    }else{
      addressLanguage.zh_CN = ''
    }
  }
}

onMounted(async() =>  {
  loading.value = true;
  await loadOptions();
  await loadTypeList()
  await loadPropertyList()
  if(params.id){
    await load()
  }
  loading.value = false;
});

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


