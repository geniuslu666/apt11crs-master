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
          :title="formValue.id > 0 ? '复制服务 #' + formValue.id : '添加服务'"
        >
          <n-tabs type="line" animated v-model:value="tabValue">
            <n-tab-pane name="base" tab="基础信息">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="服务套餐名称_简体中文" path="name_zh" :show-require-mark="true">
                    <n-input placeholder="请输入服务套餐名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" style="width: 250px" />
                  </n-form-item>
                  <n-form-item label="服务套餐名称_English" path="name_en">
                    <n-input placeholder="请输入服务套餐名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" style="width: 250px" />
                  </n-form-item>
                  <n-form-item label="服务套餐名称_日本语" path="name_ja">
                    <n-input placeholder="请输入服务套餐名称" v-model:value="nameLanguage.ja" style="width: 250px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="服务简介_简体中文" path="sub_name_zh" :show-require-mark="true">
                    <n-input type="textarea" placeholder="服务简介" v-model:value="subNameLanguage.zh" @blur="translate('sub_name','zh',subNameLanguage.zh)" style="width: 400px"/>
                  </n-form-item>
                  <n-form-item label="服务简介_English" path="sub_name_en">
                    <n-input type="textarea" placeholder="服务简介" v-model:value="subNameLanguage.en" @blur="translate('sub_name','zh',subNameLanguage.en)" style="width: 400px"/>
                  </n-form-item>
                  <n-form-item label="服务简介_日本语" path="sub_name_ja">
                    <n-input type="textarea" placeholder="服务简介" v-model:value="subNameLanguage.ja" style="width: 400px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="所属服务商" path="ispId"  :show-require-mark="true">
                    <n-select
                      v-model:value="formValue.ispId"
                      :options="ispList"
                      label-field="name"
                      value-field="id"
                      style="width: 200px"
                    />
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
                          :label="item.labelName"
                        />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="服务图集" path="imagesArr" :show-require-mark="true" :show-feedback='false'>
                    <FileChooser :maxNumber="10" v-model:value="formValue.imagesArr" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 180px;">建议尺寸：686px*324px</div>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="服务渠道" path="channel" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.channel" name="channel">
                      <n-space>
                        <n-radio :value="1">
                          到店和上门
                        </n-radio>
                        <n-radio :value="2">
                          仅上门
                        </n-radio>
                        <n-radio :value="3">
                          仅到店
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1" v-if="formValue.channel != 2">
                  <n-form-item label="适用物业" path="propertyIdsArr" :show-require-mark="true">
                    <n-select
                      placeholder="请选择物业"
                      v-model:value="formValue.propertyIdsArr"
                      :options="propertyList"
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
                  <n-form-item label="是否上架" path="serviceState" :show-require-mark="true">
                    <n-radio-group v-model:value="formValue.serviceState" name="serviceState">
                      <n-radio-button
                        v-for="status in serviceStatusArr"
                        :key="status.value"
                        :value="status.value"
                        :label="status.label"
                      />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="排序" path="sort">
                    <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" />
                  </n-form-item>
                </n-gi>

              </n-grid>
            </n-tab-pane>
            <n-tab-pane name="price" tab="价格时长">
              <n-table>
                <thead>
                <tr>
                  <th width="250">项目名中文</th>
                  <th width="250">项目名日语</th>
                  <th width="250">项目名英文</th>
                  <th width="150">项目图片(160px*160px)</th>
                  <th width="200">价格</th>
                  <th width="200">时长</th>
                  <th width="250">时长</th>
                  <th width="130">操作</th>
                </tr>
                </thead>
                <tbody>
                <tr :key="index" v-for="(item, index) in priceList">
                  <td>
                    <n-input placeholder="项目名中文" v-model:value="item.nameLanguage.zh" @blur="translatePriceItem(index,item.nameLanguage.zh,'zh_CN')"/>
                  </td>
                  <td>
                    <n-input placeholder="项目名日语" v-model:value="item.nameLanguage.ja" />
                  </td>
                  <td>
                    <n-input placeholder="项目名英文" v-model:value="item.nameLanguage.en" @blur="translatePriceItem(index,item.nameLanguage.en,'ko')"/>
                  </td>
                  <td>
                    <UploadImage :maxNumber="1" v-model:value="item.image" />
                  </td>
                  <td>
                    <n-input-number placeholder="价格" :min="0" :show-button="false" v-model:value="item.price" style="width:150px">
                      <template #suffix>
                        JPY
                      </template>
                    </n-input-number>
                  </td>
                  <td>
                    <n-input-number placeholder="时长" :min="0" :precision="0" :show-button="false" v-model:value="item.duration" style="width:150px">
                      <template #suffix>
                        分钟
                      </template>
                    </n-input-number>
                  </td>
                  <td>
                    <n-button
                      v-if="index > 0"
                      @click="delPriceItem(index)"
                      text
                      tag="a"
                      href="javascript:void(0);"
                      type="error"
                    >
                      删除
                    </n-button>
                  </td>
                </tr>
                </tbody>
              </n-table>
              <n-button type="primary" style="margin-top: 20px" @click="addPriceItem">
                新增
              </n-button>
            </n-tab-pane>
            <n-tab-pane name="content" tab="服务详情">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="服务详情_简体中文" path="content_zh" style="width: 80%" :show-require-mark="true">
                    <Editor style="height: 400px;" id="content_zh" v-model:modelValue="contentLanguage.zh" />
                  </n-form-item>
                  <n-form-item label="服务详情_繁体" path="content_zh_CN" style="width: 80%">
                    <Editor style="height: 400px;" id="content_zh_CN" v-model:modelValue="contentLanguage.zh_CN" />
                  </n-form-item>
                  <n-form-item label="服务详情_English" path="content_en" style="width: 80%">
                    <Editor style="height: 400px;" id="content_en" v-model:modelValue="contentLanguage.en" />
                  </n-form-item>
                  <n-form-item label="服务详情_日本语" path="content_ja" style="width: 80%">
                    <Editor style="height: 400px;" id="content_ja" v-model:modelValue="contentLanguage.ja" />
                  </n-form-item>
                  <n-form-item label="服务详情_韩语" path="content_ko" style="width: 80%">
                    <Editor style="height: 400px;" id="content_ko" v-model:modelValue="contentLanguage.ko" />
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

import {onMounted, reactive, ref} from "vue";
import {newState, State, rules} from "@/views/spaService/model";
import {useRouter} from "vue-router";
import {View,MaxSort,Edit} from "@/api/spaService";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Text} from "@/api/translate";
import {List as labelList} from "@/api/spaLabel";
import {All as propertyAll} from "@/api/pmsProperty";
import {Dicts} from "@/api/dict/dict";
import Editor from "@/components/Editor/editor.vue";
import {jsontoobj} from "@/utils/smjcomm";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {List as ispAllList} from "@/api/spaIsp";

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
const labelsArr = ref([])
const propertyList = ref([]);
const serviceStatusArr = ref([]);
const ispList = ref([]);
const priceList = ref([
  {
    "id": 0,
    "goodsName": "",
    "image": "",
    "nameLanguage": {
      en: '',
      zh: '',
      ja: '',
      ko: '',
      zh_CN: '',
    },
    "price": null,
    "duration": null
  }
])

const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});
const subNameLanguage = reactive({
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
  }else{
    if(lang == 'zh'){
      // 简体中文
      subNameLanguage.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      subNameLanguage.ko = ''
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
      }else{
        if(lang == 'zh'){
          // 简体中文
          subNameLanguage.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          subNameLanguage.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function translatePriceItem(index,name,lang){
  let priceListArr = priceList.value;
  if(lang == 'zh_CN'){
    priceListArr[index].nameLanguage.zh_CN = ''
  }else{
    priceListArr[index].nameLanguage.ko = ''
  }
  if(name){
    if(lang == 'ko'){
      name = name.toLowerCase()
    }
    Text({
      text: name
    }).then((_res) => {
      if(lang == 'zh_CN'){
        priceListArr[index].nameLanguage.zh_CN = _res.zh_CN
      }else{
        priceListArr[index].nameLanguage.ko = _res.ko
      }
    }).catch((err) => {

    });
  }
}

function addPriceItem(){
  priceList.value.push(
    {
      "id": 0,
      "goodsName": "",
      "image": "",
      "nameLanguage": {
        en: '',
        zh: '',
        ja: '',
        ko: '',
        zh_CN: '',
      },
      "price": null,
      "duration": null
    }
  )
}

function delPriceItem(index){
  priceList.value.splice(index, 1)
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!nameLanguage.zh){
        formBtnLoading.value = false;
        message.error('服务名称请填写完整');
        tabValue.value = 'base';
        return false;
      }
      if(!subNameLanguage.zh){
        formBtnLoading.value = false;
        message.error('服务简介请填写完整');
        tabValue.value = 'base';
        return false;
      }
      if(formValue.value.imagesArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请上传图集');
        tabValue.value = 'base';
        return false;
      }
      if(formValue.value.channel != 2 && formValue.value.propertyIdsArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择物业');
        tabValue.value = 'base';
        return false;
      }
      formValue.value.id = 0;
      formValue.value.nameLanguage = nameLanguage
      formValue.value.subNameLanguage = subNameLanguage
      formValue.value.contentLanguage = contentLanguage

      formValue.value.labelIds = formValue.value.labelIdsArr.join(',')
      formValue.value.images = formValue.value.imagesArr.join(',')
      formValue.value.propertyIds = formValue.value.channel != 2 ? formValue.value.propertyIdsArr.join(',') : ''

      let priceListArr = priceList.value;
      let priceListErr = false;
      priceListArr.forEach((item) => {
        if(!item.nameLanguage.zh || !item.nameLanguage.en || !item.nameLanguage.ja ||!item.image || !item.price || !item.duration){
          priceListErr = true;
          return true;
        }
      })
      if(priceListErr){
        formBtnLoading.value = false;
        message.error('价格列表请填写完整');
        tabValue.value = 'price';
        return false;
      }
      if(!contentLanguage.zh){
        formBtnLoading.value = false;
        message.error('服务详情请填写完整');
        tabValue.value = 'content';
        return false;
      }
      formValue.value.priceList = priceListArr;

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'spaServiceIndex', params: {  } });
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
    isLanguage: true
  })
  formValue.value = detailRes
  formValue.value.labelIdsArr = detailRes.labelIds ? detailRes.labelIds.split(',').map((item)=>{
    return parseInt(item)
  }) : [];
  formValue.value.propertyIdsArr = detailRes.propertyIds ? detailRes.propertyIds.split(",").map(item => Number(item)) : [];
  formValue.value.imagesArr =  detailRes.images ? detailRes.images.split(',').map((item)=>{
    return item
  }) : [];

  if(detailRes.nameLanguage){
    const nameLanguageObj = jsontoobj(detailRes.nameLanguage);
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
  }
  if(detailRes.subNameLanguage){
    const subNameLanguageObj = jsontoobj(detailRes.subNameLanguage);
    subNameLanguage.zh = subNameLanguageObj.zh && subNameLanguageObj.zh.content ? subNameLanguageObj.zh.content : ''
    subNameLanguage.en = subNameLanguageObj.en && subNameLanguageObj.en.content ? subNameLanguageObj.en.content : ''
    subNameLanguage.ko = subNameLanguageObj.ko && subNameLanguageObj.ko.content ? subNameLanguageObj.ko.content : ''
    subNameLanguage.ja = subNameLanguageObj.ja && subNameLanguageObj.ja.content ? subNameLanguageObj.ja.content : ''
    if(subNameLanguageObj.zh_CN){
      subNameLanguage.zh_CN = subNameLanguageObj.zh_CN.content ? subNameLanguageObj.zh_CN.content : ''
    }else if(subNameLanguageObj.zh_cn){
      subNameLanguage.zh_CN = subNameLanguageObj.zh_cn.content ? subNameLanguageObj.zh_cn.content : ''
    }else{
      subNameLanguage.zh_CN = ''
    }
  }
  if(detailRes.contentLanguage){
    const contentLanguageObj = jsontoobj(detailRes.contentLanguage);
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

  let priceListArr = [];
  if(detailRes.priceList){
    detailRes.priceList.forEach((item) => {
      let priceListItem = {};
      let itemNameLanguage = {
        en: '',
        zh: '',
        ja: '',
        ko: '',
        zh_CN: '',
      };
      if(item.nameLanguage){
        let itemNameLanguageObj = jsontoobj(item.nameLanguage);
        itemNameLanguage.zh = itemNameLanguageObj.zh && itemNameLanguageObj.zh.content ? itemNameLanguageObj.zh.content : ''
        itemNameLanguage.en = itemNameLanguageObj.en && itemNameLanguageObj.en.content ? itemNameLanguageObj.en.content : ''
        itemNameLanguage.ko = itemNameLanguageObj.ko && itemNameLanguageObj.ko.content ? itemNameLanguageObj.ko.content : ''
        itemNameLanguage.ja = itemNameLanguageObj.ja && itemNameLanguageObj.ja.content ? itemNameLanguageObj.ja.content : ''
        if(itemNameLanguageObj.zh_CN){
          itemNameLanguage.zh_CN = itemNameLanguageObj.zh_CN.content ? itemNameLanguageObj.zh_CN.content : ''
        }else if(itemNameLanguageObj.zh_cn){
          itemNameLanguage.zh_CN = itemNameLanguageObj.zh_cn.content ? itemNameLanguageObj.zh_cn.content : ''
        }else{
          itemNameLanguage.zh_CN = ''
        }
      }else{
        itemNameLanguage.zh = '';
        itemNameLanguage.en = '';
        itemNameLanguage.ko = '';
        itemNameLanguage.ja = '';
        itemNameLanguage.zh_CN = '';
      }
      priceListItem = {
        "id": item.id,
        "goodsName": item.goodsName,
        "image": item.image,
        "nameLanguage": itemNameLanguage,
        "price": item.price,
        "duration": item.duration
      }
      priceListArr.push(priceListItem)
    })
    priceList.value = priceListArr
  }
}

async function loadLabelList(){
  let labelsArrList = await labelList({
    status: 1,
    Pagination: false
  })
  labelsArr.value = labelsArrList.list
}

async function loadPropertyList(){
  let propertyArrList = await propertyAll({
    isFindClose: true
  })
  propertyList.value = propertyArrList.list
}

async function loadDict(){
  let dictsList = await Dicts({
    types: ['service_status'],
  })
  serviceStatusArr.value = dictsList.service_status
}

async function loadMax(){
  let max = await MaxSort()
  formValue.value.sort = max.sort
}

async function init(){
  formValue.value = newState(null);
}

async function loadIspList(){
  let ispArrListOrg = await ispAllList({
    status: 1,
    Pagination: false
  })
  ispList.value = ispArrListOrg.list
}

onMounted(async() => {
  show.value = true;
  await loadLabelList()
  await loadPropertyList()
  await loadDict()
  await loadIspList();
  if(parseInt(params.id) > 0){
    await load();
  }else{
    await init();
    await loadMax()
  }

  show.value = false;
});
</script>

<style lang="less"></style>


