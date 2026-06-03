<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑标签 #' + formValue.id : '添加标签'"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
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
                <n-form-item label="标签名称" path="labelName">
                  <n-input placeholder="请输入标签名称" v-model:value="formValue.labelName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="标签内容_简体中文" path="content_zh" :show-require-mark="true">
                  <n-input placeholder="请输入标签内容" v-model:value="contentLanguage.zh" @blur="translate('zh',contentLanguage.zh)" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="标签内容_English" path="content_en">
                  <n-input placeholder="请输入标签内容" v-model:value="contentLanguage.en" @blur="translate('en',contentLanguage.en)" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="标签内容_日本语" path="content_ja">
                  <n-input placeholder="请输入标签内容" v-model:value="contentLanguage.ja" />
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
                <n-form-item label="排序" path="sort">
                  <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" style="width: 100px"/>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue';
import { Edit, View, MaxSort } from '@/api/spaLabel';
import { options, State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const contentLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

function translate(lang,text){
  if(lang == 'zh'){
    // 简体中文
    contentLanguage.zh_CN = ''
  }else if(lang == 'en'){
    // 韩语
    contentLanguage.ko = ''
  }
  if(text){
    if(lang == 'en'){
      text = text.toLowerCase()
    }
    Text({
      text: text
    }).then((_res) => {
      if(lang == 'zh'){
        // 简体中文
        contentLanguage.zh_CN = _res.zh_CN
      }else if(lang == 'en'){
        // 韩语
        contentLanguage.ko = _res.ko
      }
    }).catch((err) => {

    });
  }
}

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    contentLanguage.zh = '';
    contentLanguage.en = '';
    contentLanguage.ja = '';
    contentLanguage.ko = '';
    contentLanguage.zh_CN = '';

    loading.value = true;
    MaxSort()
      .then((res) => {
        formValue.value.sort = res.sort;
      })
      .finally(() => {
        loading.value = false;
      });
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id,isLanguage: true })
    .then((res) => {
      formValue.value = res;

      if(res.contentLanguage){
        const contentLanguageObj = jsontoobj(res.contentLanguage);
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
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      formValue.value.contentLanguage = contentLanguage
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
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

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


