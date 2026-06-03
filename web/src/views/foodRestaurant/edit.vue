<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="formValue.id > 0 ? restaurantName + '-' + '餐厅设置 #' + formValue.id : '添加餐厅'"
      >
        <n-tabs type="line" animated v-model:value="tabValue" v-if="formValue.id > 0">
          <n-tab-pane name="base" tab="基础设置">
            <EditBase ref="editBaseRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="operation" tab="运营设置">
            <EditOperation ref="editOperationRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="others" tab="其他设置">
            <EditOthers ref="editOthersRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="settlement" tab="结算设置">
            <EditSettlement ref="editSettlementRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
          <n-tab-pane name="terminal" tab="终端设置">
            <EditTerminal ref="editTerminalRef" :formData="formValue" @reloadInfo="load"/>
          </n-tab-pane>
<!--          <n-tab-pane name="notice" tab="商家通知">
            <EditNotice ref="editNoticeRef" :restaurantId="params.id" @reloadInfo="load"/>
          </n-tab-pane>-->
        </n-tabs>
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
            <n-tab-pane name="base" tab="基础设置">
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
<!--                <n-gi span="1">-->
<!--                  <div style="position: relative">-->
<!--                    <div class="mt-3">-->
<!--                      <googleMap-->
<!--                        :key="mapkey"-->
<!--                        :lat="formValue.ggLat"-->
<!--                        :long="formValue.ggLng"-->
<!--                        :heightmap="260"-->
<!--                        @addressData="addressDataGLoad($event)"-->
<!--                      />-->
<!--                    </div>-->
<!--                  </div>-->
<!--                </n-gi>-->
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
                  <n-form-item label="门店抽成" path="settlementType">
                    <n-radio-group v-model:value="formValue.settlementType" name="settlementType">
                      <n-space>
                        <n-radio :value="1">
                          跟随系统
                        </n-radio>
                        <n-radio :value="2">
                          自定义
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item label="门店结算比例" path="orderConfirmHour" v-if="formValue.settlementType == 2">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.settlementRate" :show-button="false" style="width: 80px" />
                      <n-input-group-label>%</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>
            <n-tab-pane name="terminal" tab="终端设置">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi>
                  <n-form-item label="账号" path="account" :show-require-mark="true">
                    <n-input placeholder="请输入账号" v-model:value="formValue.account" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="密码" path="password" :show-require-mark="true">
                    <n-input type="password" placeholder="请输入密码" v-model:value="formValue.password" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="绑定终端" path="terminalType" :show-require-mark="true">
                    <n-button type="primary" @click="chooseTerminal('VERIFY_PRINTER')" style="margin-bottom: 25px">绑定核销打印机</n-button>
                    <n-button type="primary" @click="chooseTerminal('HAND_TERMINAL')" style="margin-bottom: 25px;margin-left: 10px">绑定手持终端</n-button>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-table>
                    <thead>
                    <tr>
                      <th>终端编号</th>
                      <th>终端名称</th>
                      <th>终端类型</th>
                      <th>品牌型号</th>
                      <th>打印联数</th>
                      <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="(item, index) of terminalArr" :key="index">
                      <td>{{ item.sn }}</td>
                      <td>{{ item.terminalName }}</td>
                      <td>
                        <n-tag
                          :type="getOptionTag(options.terminal_type, item.terminalType)"
                          size="small"
                          class="min-left-space"
                        >
                          {{ getOptionLabel(options.terminal_type, item.terminalType) }}
                        </n-tag>
                      </td>
                      <td>
                        <template v-if="item.terminalType === 'VERIFY_PRINTER'">
                          <n-tag
                            :type="getOptionTag(options.verify_brand_model, item.brandModel)"
                            size="small"
                            class="min-left-space"
                          >
                            {{ getOptionLabel(options.verify_brand_model, item.brandModel) }}
                          </n-tag>
                        </template>
                        <template v-else-if="item.terminalType === 'HAND_TERMINAL'">
                          <n-tag
                            :type="getOptionTag(options.hand_brand_model, item.brandModel)"
                            size="small"
                            class="min-left-space"
                          >
                            {{ getOptionLabel(options.hand_brand_model, item.brandModel) }}
                          </n-tag>
                        </template>
                      </td>
                      <td>
                        <template v-if="item.terminalType == 'VERIFY_PRINTER'">
                          <n-input-number v-model:value="item.printTimes" :min="1" :precision="0" :show-button="false" placeholder="打印联数" style="width: 100px"/>
                        </template>
                        <template v-else>
                          --
                        </template>
                      </td>
                      <td>
                        <n-button type="error" @click="handleDeleteTerminal(index,item.terminalType,item.id)">删除</n-button>
                      </td>
                    </tr>
                    </tbody>
                  </n-table>
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
          <ChooseTerminal ref="chooseTerminalRef" @reloadTerminal="handleChooseTerminal"/>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive} from 'vue';
import {Edit, View} from '@/api/foodRestaurant';
import { State, newState } from './model';
import {useRouter} from "vue-router";
import EditBase from "@/views/foodRestaurant/edit_base.vue";
import EditOperation from "@/views/foodRestaurant/edit_operation.vue";
import EditOthers from "@/views/foodRestaurant/edit_others.vue";
import EditSettlement from "@/views/foodRestaurant/edit_settlement.vue"
import EditTerminal from "@/views/foodRestaurant/edit_terminal.vue"
import {jsontoobj} from "@/utils/smjcomm";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {List as cuisineList} from "@/api/foodCuisine";
import {List as labelList} from "@/api/foodLabel";
import {List as typeList} from "@/api/foodCooperateType";
import {List as terminalList} from "@/api/terminal"
import {Text} from "@/api/translate";
import {List} from "@/api/foodSettlement";
import googleMap from '@/views/smjcomm/GMap.vue';
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {loadOptions as loadTerminalOptions, options} from "@/views/terminal/model";
import ChooseTerminal from "@/components/ChooseTerminal/chooseTerminal.vue";

const tabsViewStore = useTabsViewStore();
const formRef = ref<any>({});
const message = useMessage();
const tabValue = ref('base')
const newTabValue = ref('base')
const show = ref(false);
const router = useRouter();
const params = router.currentRoute.value.params;
const formValue = ref<State>(newState(null));
const editBaseRef = ref();
const editOperationRef = ref();
const editOthersRef = ref();
const editSettlementRef = ref();
const editTerminalRef = ref();
const restaurantName = ref('');
const rules = ref({})
const settingStore = useProjectSettingStore();
const cuisinesArr = ref([])
const labelsArr = ref([])
const cooperateTypeArr = ref([])
const settlementList = ref([]);
let mapkey = ref(1);
const formBtnLoading = ref(false);
const chooseTerminalRef = ref();
const chooseTerminalType = ref('')
const terminalArr = ref([])
const terminalIdsArr = ref([])

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
    }).catch((err) => {

    });
  }
}

const addressDataGLoad = (res) => {
  console.log('addressData', res);
  formValue.value.ggLat = res.lat + '';
  formValue.value.ggLng = res.lng + '';
};

function confirmForm(e) {
  formValue.value.terminalList = terminalArr.value

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
      if(!formValue.value.detailAddress){
        formBtnLoading.value = false;
        message.error('请输入详细地址');
        return false;
      }
      if(!formValue.value.settlementId){
        formBtnLoading.value = false;
        message.error('请选择结算模式');
        newTabValue.value = 'settlement'
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
          tabsViewStore.closeSignal('2');
          router.push({ name: 'foodRestaurantIndex', params: {  } });
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

async function loadSettlementList(){
  let settlementArrListOrg = await List({
    status: 1,
    Pagination: false
  })
  settlementList.value = settlementArrListOrg.list
}

function chooseTerminal(type){
  chooseTerminalType.value = type
  let chooseTerminalIds = '';
  if(type == 'VERIFY_PRINTER'){
    chooseTerminalIds = formValue.value.printTerminalIds
  }else{
    chooseTerminalIds = formValue.value.handTerminalIds
  }
  chooseTerminalRef.value.openModal('RESTAURANT', type, chooseTerminalIds, formValue.value.id);
}

async function getTerminalList(){
  let terminalIds = (formValue.value.printTerminalIds + ',' +  formValue.value.handTerminalIds).replace(/,$/, '')
  const res = await terminalList({
    pagination: false,
    terminalIds: terminalIds
  });
  if(res.list && res.list.length > 0){
    terminalArr.value = res.list;
    terminalIdsArr.value = res.list.map((item) => {
      return parseInt(item.id)
    })
    formValue.value.terminalIds = terminalIdsArr.value.join(',')
  }else{
    terminalArr.value = []
    terminalIdsArr.value = []
    formValue.value.terminalIds = ''
  }
  show.value = false;
}


function handleDeleteTerminal(index,type,targetId){
  terminalArr.value.splice(index, 1);
  terminalIdsArr.value.splice(index, 1);
  formValue.value.terminalIds = terminalIdsArr.value.join(',')
  if(type == 'VERIFY_PRINTER'){
    formValue.value.printTerminalIds = formValue.value.printTerminalIds.split(',').filter(id => parseInt(id) !== targetId).join(',')
  }else{
    formValue.value.handTerminalIds = formValue.value.handTerminalIds.split(',').filter(id => parseInt(id) !== targetId).join(',')
  }
}

async function handleChooseTerminal(record){
  if(chooseTerminalType.value == 'VERIFY_PRINTER'){
    // 打印机
    formValue.value.printTerminalIds = record
  }else{
    // 手持
    formValue.value.handTerminalIds = record
  }
  show.value = true;
  await getTerminalList()
}

async function load() {
  show.value = true;
  const res = await View({
    id: params.id,
    isLanguage: true
  });
  if(res){
    formValue.value = res;
    let nameLanguage = jsontoobj(formValue.value.nameLanguage);
    restaurantName.value = nameLanguage.zh && nameLanguage.zh.content ? nameLanguage.zh.content : ''

    let printTerminalIdsArr = []
    let handTerminalIdsArr = []
    if(res.terminalList && res.terminalList.length > 0){
      res.terminalList.forEach((item,index) => {
        if(item.terminalInfo){
          if(item.terminalInfo.terminalType == 'VERIFY_PRINTER'){
            printTerminalIdsArr.push(parseInt(item.terminalId))
          }else{
            handTerminalIdsArr.push(parseInt(item.terminalId))
          }
        }
      })
    }
    formValue.value.printTerminalIds = printTerminalIdsArr.join(',')
    formValue.value.handTerminalIds = handTerminalIdsArr.join(',')
  }
  show.value = false;
}

onMounted(async() => {
  loadTerminalOptions()
  if(params.id > 0){
    await load();
  }else{
    show.value = true;
    await loadCuisineList()
    await loadLabelList()
    await loadTypeList()
    await loadSettlementList()
    show.value = false;
  }
});
</script>

<style lang="less"></style>


