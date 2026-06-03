<template>
  <div>
    <n-drawer v-model:show="showModal" :close-on-esc="false" :width="dialogWidth" @after-leave="closeForm">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑门店 #' + formValue.id : '添加门店' }}</div>
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
            label-placement="top"
            label-width="auto"
          >
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                门店信息
              </div>
              <n-grid :cols="1" x-gap="80">
                <n-gi>
                  <n-form-item v-if="initMchId <= 0" label="所属商户" path="mchId" :show-require-mark="true">
                    <n-select
                      v-model:value="formValue.mchId"
                      :options="mchList"
                      clearable
                      filterable
                      label-field="name"
                      value-field="id"
                      style="width: 320px"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="门店名称_简体中文" path="name_zh" :show-require-mark="true">
                    <n-input placeholder="简体中文门店名称" v-model:value="nameLanguage.zh" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="门店名称_日本语" path="name_ja" :show-require-mark="true">
                    <n-input placeholder="日本语门店名称" v-model:value="nameLanguage.ja" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="门店名称_English" path="name_en" :show-require-mark="true">
                    <n-input placeholder="英语门店名称" v-model:value="nameLanguage.en" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="图集" path="images">
                    <FileChooser :maxNumber="10" v-model:value="formValue.imagesArr" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="门店电话" path="phone">
                    <n-input placeholder="请输入门店电话" v-model:value="formValue.phone" :style="{ width: '300px' }" />
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
                  <n-form-item label="门店地址" path="detailAddress" :show-require-mark="true">
                    <n-input placeholder="请输入详细地址" v-model:value="formValue.detailAddress" style="width: 300px"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi span="1">-->
<!--                  <div style="position: relative;margin-bottom: 24px">-->
<!--                    <div>-->
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
                <n-gi>
                  <n-form-item label="账号" path="account" :show-require-mark="true">
                    <n-input placeholder="请输入账号" v-model:value="formValue.account" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="密码" path="password" :show-require-mark="true">
                    <n-input type="password" :placeholder="formValue.id > 0 ? '不填则不修改' : '请输入密码'" v-model:value="formValue.password" :style="{ width: '300px' }" />
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
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <ChooseTerminal ref="chooseTerminalRef" @reloadTerminal="handleChooseTerminal"/>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed} from 'vue';
import { View, Edit } from '@/api/thMchStore';
import { State, newState } from './model';
import {options} from '@/views/terminal/model';
import { useMessage } from 'naive-ui';
import {adaModalWidth, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {All} from "@/api/thMch";
import {List} from "@/api/terminal"
import {Text} from "@/api/translate";
import googleMap from '@/views/smjcomm/GMap.vue';
import ChooseTerminal from "@/components/ChooseTerminal/chooseTerminal.vue";
import {jsontoobj} from "@/utils/smjcomm";

const initMchId = ref(0)
const emit = defineEmits(['reloadTable','closeEdit']);
const chooseTerminalRef = ref();
let mapkey = ref(1);
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(950);
});
const chooseTerminalType = ref('')
const terminalArr = ref([])
const message = useMessage();
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const formValue = ref<State>(newState(null));
const mchList = ref([])
const terminalIdsArr = ref([])
const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const addressDataGLoad = (res) => {
  formValue.value.ggLat = res.lat + '';
  formValue.value.ggLng = res.lng + '';
};

function translate(type,lang,text){
  if(type == 'name'){
    if(lang == 'zh'){
      // 简体中文
      nameLanguage.value.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.value.ko = ''
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
          nameLanguage.value.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          nameLanguage.value.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function chooseTerminal(type){
  chooseTerminalType.value = type
  let chooseTerminalIds = '';
  if(type == 'VERIFY_PRINTER'){
    chooseTerminalIds = formValue.value.printTerminalIds
  }else{
    chooseTerminalIds = formValue.value.handTerminalIds
  }
  chooseTerminalRef.value.openModal('TH_COUPON', type, chooseTerminalIds, formValue.value.id);
}

async function getAllMch(){
  const res = await All({});
  mchList.value = res.list;
}

async function getTerminalList(){
  let terminalIds = (formValue.value.printTerminalIds + ',' +  formValue.value.handTerminalIds).replace(/,$/, '')
  const res = await List({
    pagination: false,
    terminalIds: terminalIds,
    fromStore: 1,
    NeedPrintTimesStoreId: formValue.value.id
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
  loading.value = false;
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

function confirmForm(e) {
  formValue.value.terminalList = terminalArr.value
  // 门店名称-多语言
  nameLanguage.value.zh_CN = nameLanguage.value.ja
  nameLanguage.value.ko = nameLanguage.value.ja
  formValue.value.nameLanguage = nameLanguage.value

  formValue.value.storeName = ""

  if(!formValue.value.mchId){
    formBtnLoading.value = false;
    message.error('请选择所属商户');
    return false;
  }

  if(!nameLanguage.value.zh || !nameLanguage.value.en || !nameLanguage.value.ja){
    formBtnLoading.value = false;
    message.error('门店名称请填写完整');
    return false;
  }

/*  if(!formValue.value.phone){
    formBtnLoading.value = false;
    message.error('请输入门店电话');
    return false;
  }*/

  if(!formValue.value.detailAddress){
    formBtnLoading.value = false;
    message.error('请输入门店地址');
    return false;
  }

  if(!formValue.value.ggLat || !formValue.value.ggLng){
    formBtnLoading.value = false;
    message.error('请选择门店定位');
    return false;
  }

  if(!formValue.value.account){
    formBtnLoading.value = false;
    message.error('请输入账号');
    return false;
  }

  if(formValue.value.id <= 0 && !formValue.value.password){
    formBtnLoading.value = false;
    message.error('请输入密码');
    return false;
  }

/*  if(!formValue.value.terminalIds){
    formBtnLoading.value = false;
    message.error('请选择绑定终端');
    return false;
  }*/
  // return

  formValue.value.images = formValue.value.imagesArr.join(',')

  e.preventDefault();

  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

async function getInfo(id){
  const res = await View({ id: id, isLanguage: true });
  formValue.value = res;
  formValue.value.password = ''

  if(res.nameLanguage){
    const nameLanguageObj = jsontoobj(res.nameLanguage);
    nameLanguage.value.zh = nameLanguageObj.zh && nameLanguageObj.zh.content ? nameLanguageObj.zh.content : ''
    nameLanguage.value.en = nameLanguageObj.en && nameLanguageObj.en.content ? nameLanguageObj.en.content : ''
    nameLanguage.value.ko = nameLanguageObj.ko && nameLanguageObj.ko.content ? nameLanguageObj.ko.content : ''
    nameLanguage.value.ja = nameLanguageObj.ja && nameLanguageObj.ja.content ? nameLanguageObj.ja.content : ''
    if(nameLanguageObj.zh_CN){
      nameLanguage.value.zh_CN = nameLanguageObj.zh_CN.content ? nameLanguageObj.zh_CN.content : ''
    }else if(nameLanguageObj.zh_cn){
      nameLanguage.value.zh_CN = nameLanguageObj.zh_cn.content ? nameLanguageObj.zh_cn.content : ''
    }else{
      nameLanguage.value.zh_CN = ''
    }
  }

  formValue.value.imagesArr = res.images ? res.images.split(',').map((item)=>{
    return item
  }) : [];

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

async function handleChooseTerminal(record){
  if(chooseTerminalType.value == 'VERIFY_PRINTER'){
    // 打印机
    formValue.value.printTerminalIds = record
  }else{
    // 手持
    formValue.value.handTerminalIds = record
  }
  loading.value = true;
  await getTerminalList()
}

async function openModal(state: State, mchId) {
  showModal.value = true;
  loading.value = true;

  await getAllMch();

  // 新增
  if (!state || state.id < 1) {
    terminalArr.value = []
    formValue.value = newState(state);

    if(parseInt(mchId) > 0){
      initMchId.value = parseInt(mchId)
      formValue.value.mchId = parseInt(mchId)
    }

    nameLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };

    loading.value = false;

    return;
  }

  await getInfo(state.id)
  await getTerminalList()
  loading.value = false;

}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  emit('closeEdit');
}

defineExpose({
  openModal,
});
</script>

<style lang="less">
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


