<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑餐厅菜系 #' + formValue.id : '添加餐厅菜系' }}</div>
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
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="菜系名称-简体中文" path="cuisineName_zh" :show-require-mark="true">
                  <n-input placeholder="请输入菜系名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="菜系名称_English" path="cuisineName_en">
                  <n-input placeholder="请输入菜系名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="菜系名称_日本语" path="cuisineName_ja">
                  <n-input placeholder="请输入菜系名称" v-model:value="nameLanguage.ja" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="图片" path="pic">
                  <FileChooser1 v-model:value="formValue.pic" :maxNumber="1"
                                fileType="default"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="状态" path="status">
                  <a-radio-group v-model:value="formValue.status" name="status">
                    <a-radio v-for="model in options.sys_normal_disable"
                             :key="model.value"
                             :value="model.value">{{model.label}}</a-radio>
                  </a-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="排序" path="sort">
                  <n-input-number placeholder="请输入排序" v-model:value="formValue.sort"/>
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
import { Edit, View, MaxSort } from '@/api/foodCuisine';
import { options, State, newState, rules } from './model';
import { useMessage } from 'naive-ui';
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";
import {adaModalWidth} from "@/utils/hotgo";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

function translate(type,lang,text){
  if(lang == 'zh'){
    // 简体中文
    nameLanguage.zh_CN = ''
  }else if(lang == 'en'){
    // 韩语
    nameLanguage.ko = ''
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
        nameLanguage.zh_CN = _res.zh_CN
      }else if(lang == 'en'){
        // 韩语
        nameLanguage.ko = _res.ko
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
    nameLanguage.zh = '';
    nameLanguage.en = '';
    nameLanguage.ja = '';
    nameLanguage.ko = '';
    nameLanguage.zh_CN = '';

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

      if(res.nameLanguage){
        const nameLanguageObj = jsontoobj(res.nameLanguage);
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
      formValue.value.nameLanguage = nameLanguage
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


