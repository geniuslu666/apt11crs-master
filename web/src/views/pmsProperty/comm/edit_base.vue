<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" @mask-click="closeForm">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ drawTitle }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="submitBtn" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            v-if="editType != 'TagList'"
            ref="formRef"
            :model="formValue"
            label-placement="top"
            label-width="auto"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1" v-if="editType=='Address'">
                <n-form-item label="所属地区" path="regionId" :show-require-mark="false">
                  <n-select
                    v-model:value="propertyRegionId"
                    :options="regionList"
                    clearable
                    filterable
                    label-field="name"
                    value-field="id"
                    style="width: 300px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="简体中文" path="zh" :show-feedback="false">
                  <n-input type="textarea" placeholder="简体中文" v-model:value="formValue.zh"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1" style="margin-top: 25px">
                <n-form-item label="语言同步" path="zh" :show-feedback="false">
                  <n-checkbox-group v-model:value="transItem" @update:value="handleTransItem">
                    <n-space item-style="display: flex;">
                      <n-checkbox value="ja" label="日语"/>
                      <n-checkbox value="en" label="英语"/>
                      <n-checkbox value="ko" label="韩语"/>
                      <n-checkbox value="zh_CN" label="繁体中文"/>
                    </n-space>
                  </n-checkbox-group>
                  <n-button :loading="transBtnLoading" @click="handleTrans" type="primary" ghost
                            size="small">
                    <img src="@/assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                    一键翻译
                  </n-button>
                </n-form-item>
              </n-gi>
              <n-gi span="1" style="margin-top: 25px">
                <n-form-item label="日本语" path="ja">
                  <n-input type="textarea" placeholder="日本语" v-model:value="formValue.ja"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="English" path="en">
                  <n-input type="textarea" placeholder="English" v-model:value="formValue.en"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="한국어" path="ko">
                  <n-input type="textarea" placeholder="한국어" v-model:value="formValue.ko"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="繁体中文" path="zh_CN">
                  <n-input type="textarea" placeholder="繁体中文" v-model:value="formValue.zh_CN"/>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>


          <n-form
            v-else
            ref="formRef"
            :model="formValueTag"
            label-placement="top"
            label-width="auto"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="简体中文" path="zh">
                  <n-dynamic-tags v-model:value="formValueTag.zh"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="日本语" path="ja">
                  <n-dynamic-tags v-model:value="formValueTag.ja"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="English" path="en">
                  <n-dynamic-tags v-model:value="formValueTag.en"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="한국어" path="ko">
                  <n-dynamic-tags v-model:value="formValueTag.ko"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="繁体中文" path="zh_CN">
                  <n-dynamic-tags v-model:value="formValueTag.zh_CN"/>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, reactive, ref} from 'vue';
import {editPropertyLanguage} from '@/api/pmsProperty';
import {List} from '@/api/pmsPropertyRegion';
import {Text} from '@/api/translate';
import {useDialog, useMessage} from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";
import {pmsCancelRateedit} from "@/api/comm";

const dialog = useDialog();
const emit = defineEmits(['reloadTable']);
const drawTitle = ref('');
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(620);
});
const editType = ref('');
const pid = ref(0);
const formValue = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
})
const transItem = ref(['en', 'ja', 'ko', 'zh', 'zh_CN']);
let formValueTag = reactive({
  en: [],
  zh: [],
  ja: [],
  ko: [],
  zh_CN: [],
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const transBtnLoading = ref(false);
const regionList = ref([]);
const propertyRegionId = ref(0);

async function openModal(id, type, title, data, regionId) {
  pid.value = id;
  editType.value = type;
  drawTitle.value = title;

  showModal.value = true;
  loading.value = true;

  await getAllRegion();

  propertyRegionId.value = regionId;

  // console.log('data',regionId)
  if (type == 'TagList') {
    formValueTag.en = data.en ? JSON.parse(data.en.content) : [];
    formValueTag.zh = data.zh ? JSON.parse(data.zh.content) : [];
    formValueTag.ja = data.ja ? JSON.parse(data.ja.content) : [];
    formValueTag.ko = data.ko ? JSON.parse(data.ko.content) : [];
    formValueTag.zh_CN = data.zh_CN ? JSON.parse(data.zh_CN.content) : [];
  } else {
    formValue.value.zh = data.zh && data.zh.content ? data.zh.content : ''
    formValue.value.en = data.en && data.en.content ? data.en.content : ''
    formValue.value.ko = data.ko && data.ko.content ? data.ko.content : ''
    formValue.value.ja = data.ja && data.ja.content ? data.ja.content : ''
    formValue.value.zh_CN = data.zh_CN && data.zh_CN.content ? data.zh_CN.content : ''
  }
  loading.value = false;

}

function submitBtn(e){
  e.preventDefault();
  if(editType.value == 'CancelRateeditName'){
    CancelRateeditConfirmForm()
  }else{
    confirmForm()
  }
}

function CancelRateeditConfirmForm() {
  console.log(222)
  formBtnLoading.value = true;
  let dataobj = {
    id: pid.value,
    nameLanguage: formValue.value
  }

  pmsCancelRateedit(dataobj).then((_res) => {
    message.success('操作成功');
    setTimeout(() => {
      closeForm();
      formBtnLoading.value = false;
      emit('reloadTable');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function confirmForm() {
  console.log(111)
  formBtnLoading.value = true;

  let dataobj = {
    type: editType.value,
    id: pid.value,
    Data: {
      name: {},
      description: {},
      address: {},
      tag_list: {},
      room_des: {},
      surroundings: {},
      bus_station: {},
      subway: {},
      required_book: {},
      check_in_guide: {},
      cancel_policy: {},
    },
    regionId: regionList.value.find(item => item.id == propertyRegionId.value)?.id ,
  };
  if (editType.value == 'Name') {
    dataobj.Data.name = formValue.value;
  } else if (editType.value == 'Address') {
    dataobj.Data.address = formValue.value;
  } else if (editType.value == 'Description') {
    dataobj.Data.description = formValue.value;
  } else if (editType.value == 'BusStation') {
    dataobj.Data.bus_station = formValue.value;
  } else if (editType.value == 'RequiredBook') {
    dataobj.Data.required_book = formValue.value;
  } else if (editType.value == 'CheckInGuide') {
    dataobj.Data.check_in_guide = formValue.value;
  } else if (editType.value == 'TagList') {
    const aa = JSON.parse(JSON.stringify(formValueTag));
    aa.en = JSON.stringify(aa.en);
    aa.zh = JSON.stringify(aa.zh);
    aa.ja = JSON.stringify(aa.ja);
    aa.ko = JSON.stringify(aa.ko);
    aa.zh_CN = JSON.stringify(aa.zh_CN);
    dataobj.Data.tag_list = aa;
  }

  editPropertyLanguage(dataobj).then((_res) => {
    message.success('操作成功');
    setTimeout(() => {
      closeForm();
      formBtnLoading.value = false;
      emit('reloadTable');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  transItem.value = [];
  // pid.value = 0;
  // editType.value = '';
}

function handleTransItem(value) {
  console.log(value)
  transItem.value = value
}

function handleTrans() {
  transBtnLoading.value = true;
  if (!formValue.value.zh) {
    message.error('请先填写简体中文的内容');
    transBtnLoading.value = false;
    return false;
  }
  if (transItem.value.length <= 0) {
    message.error('请选择要进行翻译的语言');
    transBtnLoading.value = false;
    return false;
  }

  dialog.warning({
    title: '警告',
    content: '翻译会覆盖已选语言下所填写的内容，你确定要进行翻译吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Text({
        text: formValue.value.zh
      }).then((_res) => {
        message.success('翻译成功');
        if (transItem.value.indexOf('ja') !== -1) {
          formValue.value.ja = _res.ja
        }
        if (transItem.value.indexOf('ko') !== -1) {
          formValue.value.ko = _res.ko
        }
        if (transItem.value.indexOf('en') !== -1) {
          formValue.value.en = _res.en
        }
        if (transItem.value.indexOf('zh_CN') !== -1) {
          formValue.value.zh_CN = _res.zh_CN
        }
        setTimeout(() => {
          transBtnLoading.value = false;
        });
      }).catch((err) => {
        message.error('翻译失败');
        transBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      transBtnLoading.value = false;
    }
  });

}

async function getAllRegion(){
  const res = await List({
    pagination: false
  });
  regionList.value = res.list;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>
