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
            <n-form-item label="餐厅名称" path="name_zh" :show-require-mark="true">
              <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.zh" style="width: 300px" />
            </n-form-item>
          </n-gi>
<!--          <n-gi span="1">
            <n-form-item label="餐厅名称_English" path="name_en">
              <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" style="width: 35%" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="餐厅名称_日本语" path="name_ja">
              <n-input placeholder="请输入餐厅名称" v-model:value="nameLanguage.ja" style="width: 35%" />
            </n-form-item>
          </n-gi>-->
          <n-gi span="1">
            <n-form-item label="菜系" path="cuisineIdsArr">
              <n-checkbox-group v-model:value="formValue.cuisineIdsArr">
                <n-space>
                  <n-checkbox
                    v-for="item in cuisinesArr"
                    :key="item.id"
                    :value="item.id"
                    :label="item.cuisineName"
                  />
                </n-space>
              </n-checkbox-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="标签" path="labelIdsArr">
              <n-checkbox-group v-model:value="formValue.labelIdsArr">
                <n-space>
                  <n-checkbox
                    v-for="item in labelsArr"
                    :key="item.id"
                    :value="item.id"
                    :label="item.name"
                  />
                </n-space>
              </n-checkbox-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="营业类型" path="cooperateTypeId" :show-require-mark="true">
              <n-radio-group v-model:value="formValue.cooperateTypeId">
                <n-radio-button
                  v-for="item in cooperateTypeArr"
                  :key="item.id"
                  :value="item.id"
                  :label="item.typeName"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1" v-if="formValue.cooperateTypeId == 4">
            <n-form-item label="Toreta餐厅Id" path="phone">
              <n-input placeholder="请输入toretaId" v-model:value="formValue.toretaId" style="width: 400px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="LOGO" path="logo" :show-feedback='false'>
              <UploadImage :maxNumber="1" v-model:value="formValue.logo" />
            </n-form-item>
            <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：638px*300px</div>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="图集" path="imagesArr" :show-require-mark="true" :show-feedback='false'>
              <FileChooser :maxNumber="10" v-model:value="formValue.imagesArr" />
            </n-form-item>
            <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：638px*300px</div>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="简介_简体中文" path="content_zh">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="contentLanguage.zh" @blur="translate('content','zh',contentLanguage.zh)" style="width: 500px" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="简介_English" path="content_en">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="contentLanguage.en" @blur="translate('content','en',contentLanguage.en)" style="width: 500px" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="简介_日本语" path="content_ja">
              <n-input type="textarea" placeholder="请输入简介" v-model:value="contentLanguage.ja" style="width: 500px" />
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="联系电话" path="phone">
              <n-input placeholder="请输入联系电话" v-model:value="formValue.phone" style="width: 150px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="营业时间" path="openTime">
              <n-input placeholder="请输入营业时间" v-model:value="formValue.openTime" style="width: 200px"/>
            </n-form-item>
          </n-gi>
<!--          <n-gi span="1">
            <n-form-item label="餐厅地址" path="areaId" :show-require-mark="true">
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
          </n-gi>-->
          <n-gi span="1">
            <n-form-item label="餐厅纬度" path="ggLat" :show-require-mark="true">
              <n-input placeholder="请输入餐厅纬度" v-model:value="formValue.ggLat" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="餐厅经度" path="ggLng" :show-require-mark="true">
              <n-input placeholder="请输入餐厅经度" v-model:value="formValue.ggLng" style="width: 250px"/>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="餐厅地址" path="detailAddress">
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
        <div style="text-align: center;margin-top: 30px">
          <n-space justify="center">
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
              确定
            </n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>
<script setup lang="ts">
import {onMounted, reactive, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {Edit} from "@/api/foodRestaurant";
import {useMessage,TreeSelectOverrideNodeClickBehavior} from "naive-ui";
import {List as cuisineList} from "@/api/foodCuisine";
import {List as labelList} from "@/api/foodLabel";
import {List as typeList} from "@/api/foodCooperateType";
import {List as areaList} from "@/api/foodArea";
import googleMap from '@/views/smjcomm/GMap.vue';
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {jsontoobj} from "@/utils/smjcomm";
import {Text} from "@/api/translate";

const props = defineProps({
  formData: {
    type: Object || null,
    default: null,
  },
});
const emit = defineEmits(['reloadInfo']);
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const show = ref(false);
const formValue = ref({
  type: 'basic',
  id: 0,
  toretaId: '',
  nameLanguage: null,
  contentLanguage: null,
  cuisineIds: '',
  cuisineIdsArr: [],
  labelIds: '',
  labelIdsArr: [],
  cooperateTypeId: null,
  logo: '',
  imagesArr: [],
  images: '',
  content: '',
  phone: '',
  openTime: '',
  areaId: null,
  detailAddress: '',
  ggLat: '34.67100087743556',
  ggLng: '135.49982492658245',
  lat: '39.923678',
  lng: '116.403478'
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();
const cuisinesArr = ref([])
const labelsArr = ref([])
const cooperateTypeArr = ref([])
const areaArr = ref([])
const override: TreeSelectOverrideNodeClickBehavior = ({ option }) => {
  if (option.children) {
    return 'toggleExpand'
  }
  return 'default'
}
let mapkey = ref(1);

const rules = ref({
  cuisineIdsArr: {
    type: 'array',
    required: true,
    trigger: 'change',
    message: '请选择菜系'
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入简介'
  },
});
const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const contentLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.ko = ''
    }
  }
  if(type == 'content'){
    if(lang == 'zh'){
      // 简体中文
      contentLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      contentLanguage.ko = ''
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
          nameLanguage.ko = _res.ko
        }
      }
      if(type == 'content'){
        if(lang == 'zh'){
          // 简体中文
          contentLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          contentLanguage.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.cooperateTypeId){
        formBtnLoading.value = false;
        message.error('请选择营业类型');
        return false;
      }
      if(formValue.value.imagesArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请上传图集');
        return false;
      }
      /*if(!formValue.value.areaId){
        formBtnLoading.value = false;
        message.error('请选择餐厅地址');
        return false;
      }*/
      if(!formValue.value.detailAddress){
        formBtnLoading.value = false;
        message.error('请输入详细地址');
        return false;
      }

      nameLanguage.en = nameLanguage.zh
      nameLanguage.ja = nameLanguage.zh
      nameLanguage.ko = nameLanguage.zh
      nameLanguage.zh_CN = nameLanguage.zh
      formValue.value.nameLanguage = nameLanguage
      formValue.value.contentLanguage = contentLanguage
      formValue.value.cuisineIds = formValue.value.cuisineIdsArr.join(',')
      formValue.value.labelIds = formValue.value.labelIdsArr.join(',')
      formValue.value.images = formValue.value.imagesArr.join(',')
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
            emit('reloadInfo');
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'foodRestaurantIndex', params: {  } });
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

const addressDataGLoad = (res) => {
  console.log('addressData', res);
  formValue.value.ggLat = res.lat + '';
  formValue.value.ggLng = res.lng + '';
};

async function loadCuisineList(){
  const cuisinesArrList = await cuisineList({
    status: 1,
    Pagination: false
  })
  cuisinesArr.value = cuisinesArrList.list
}

async function loadLabelList(){
  const labelsArrList = await labelList({
    type: 1,
    status: 1,
    Pagination: false
  })
  labelsArr.value = labelsArrList.list
}

async function loadTypeList(){
  const cooperateTypeArrList = await typeList({
    status: 1,
    Pagination: false
  })
  cooperateTypeArr.value = cooperateTypeArrList.list
}
async function loadAreaList(){
  const areaArrList = await areaList({
    status: 1,
    Pagination: false
  })
  areaArr.value = areaArrList.list
}

async function initData(data){
  formValue.value = {
    type: 'basic',
    id: data.id,
    toretaId: data.toretaId,
    nameLanguage: data.nameLanguage,
    contentLanguage: data.contentLanguage,
    cuisineIds: data.cuisineIds,
    cuisineIdsArr: data.cuisineIds ? data.cuisineIds.split(',').map((item)=>{
      return parseInt(item)
    }) : [],
    labelIds: data.labelIds,
    labelIdsArr: data.labelIds ? data.labelIds.split(',').map((item)=>{
      return parseInt(item)
    }) : [],
    cooperateTypeId: data.cooperateTypeId,
    logo: data.logo,
    images: data.images,
    imagesArr: data.images ? data.images.split(',').map((item)=>{
      return item
    }) : [],
    content: data.content,
    phone: data.phone,
    openTime: data.openTime,
    areaId: data.areaId,
    detailAddress: data.detailAddress,
    ggLat: data.ggLat,
    ggLng: data.ggLng,
    lat: data.lat,
    lng: data.lng,
  }
  if(data.nameLanguage){
    const nameLanguageObj = jsontoobj(data.nameLanguage);
    nameLanguage.zh = nameLanguageObj.zh && nameLanguageObj.zh.content ? nameLanguageObj.zh.content : ''
    nameLanguage.en = nameLanguageObj.en && nameLanguageObj.en.content ? nameLanguageObj.en.content : ''
    nameLanguage.ko = nameLanguageObj.ko && nameLanguageObj.ko.content ? nameLanguageObj.ko.content : ''
    nameLanguage.ja = nameLanguageObj.ja && nameLanguageObj.ja.content ? nameLanguageObj.ja.content : ''
    if(nameLanguageObj.zh_CN){
      nameLanguage.zh_CN = nameLanguageObj.zh_CN.content ? nameLanguageObj.zh_CN.content : ''
    }else if(nameLanguageObj.zh_cn){
      nameLanguage.zh_CN = nameLanguageObj.zh_cn.content ? nameLanguageObj.zh_cn.content : ''
    }else{
      nameLanguage.zh_CN = ''
    }
    console.log("nameLanguage", nameLanguage);
  }
  if(data.contentLanguage){
    const contentLanguageObj = jsontoobj(data.contentLanguage);
    contentLanguage.zh = contentLanguageObj.zh && contentLanguageObj.zh.content ? contentLanguageObj.zh.content : ''
    contentLanguage.en = contentLanguageObj.en && contentLanguageObj.en.content ? contentLanguageObj.en.content : ''
    contentLanguage.ko = contentLanguageObj.ko && contentLanguageObj.ko.content ? contentLanguageObj.ko.content : ''
    contentLanguage.ja = contentLanguageObj.ja && contentLanguageObj.ja.content ? contentLanguageObj.ja.content : ''
    if(contentLanguageObj.zh_CN){
      contentLanguage.zh_CN = contentLanguageObj.zh_CN.content ? contentLanguageObj.zh_CN.content : ''
    }else if(contentLanguageObj.zh_cn){
      contentLanguage.zh_CN = contentLanguageObj.zh_cn.content ? contentLanguageObj.zh_cn.content : ''
    }else{
      contentLanguage.zh_CN = ''
    }
  }
}

watch(
  () => props.formData,
  async(value) => {
    await initData(value)
  }
);

onMounted(async() => {
  show.value = true;
  await loadCuisineList()
  await loadLabelList()
  await loadTypeList()
  await loadAreaList()
  await initData(props.formData)

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
