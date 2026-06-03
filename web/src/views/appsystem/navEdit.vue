<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑导航 #' + formValue.id : '添加导航' }}</div>
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
          <n-form :model="formValue" ref="formRef"  label-placement="top"
                  label-width="auto"
                  require-mark-placement="right-hanging"
          >
            <div class="level-detail-div">
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="导航名称_简体中文" path="name_zh" :show-require-mark="true">
                    <n-input placeholder="简体中文导航名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="导航名称_日本语" path="name_ja" :show-require-mark="true">
                    <n-input placeholder="日本语导航名称" v-model:value="nameLanguage.ja" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="导航名称_English" path="name_en" :show-require-mark="true">
                    <n-input placeholder="英语导航名称" v-model:value="nameLanguage.en" @blur="translate('name','en',nameLanguage.en)" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>

                <n-gi>
                  <n-form-item label="标签_简体中文" path="tag_zh">
                    <n-input placeholder="简体中文标签" v-model:value="tagLanguage.zh" @blur="translate('tag','zh',tagLanguage.zh)" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="标签_日本语" path="tag_ja">
                    <n-input placeholder="日本语标签" v-model:value="tagLanguage.ja" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="标签_English" path="tag_en">
                    <n-input placeholder="英语标签" v-model:value="tagLanguage.en" @blur="translate('tag','en',tagLanguage.en)" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>

                <n-gi>
                  <n-form-item label="图标" path="image" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.image" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px">建议尺寸：56px*56px</div>
                </n-gi>
                <n-gi>
                  <n-form-item label="内外联" path="chain" :show-require-mark="true">
                    <a-radio-group v-model:value="formValue.chain" name="chain">
                      <a-radio v-for="model in options.chain"
                               :key="model.value"
                               :value="model.value">{{model.label}}</a-radio>
                    </a-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.chain == 'OUT'">
                  <n-form-item label="外链跳转方式" path="linkOpenType">
                    <n-radio-group v-model:value="formValue.linkOpenType" name="linkOpenType">
                      <n-space>
                        <n-radio :value="1">
                          内部webview
                        </n-radio>
                        <n-radio :value="2">
                          外部浏览器
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.chain == 'IN'">
                  <n-form-item label="App链接" path="appLink" :show-require-mark="true">
                    <n-input placeholder="请输入App链接" v-model:value="formValue.appLink" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.chain == 'IN'">
                  <n-form-item label="小程序链接" path="wxLink" :show-require-mark="true">
                    <n-input placeholder="请输入小程序链接" v-model:value="formValue.wxLink" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.chain == 'OUT'">
                  <n-form-item label="链接内容" path="path" :show-require-mark="true">
                    <n-input placeholder="请输入链接内容" v-model:value="formValue.appLink" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="小程序显示" path="status">
                    <n-input-group style="margin-left: 8px">
                      <n-switch v-model:value="formValue.minappStatus" :unchecked-value="2" :checked-value="1">
                        <template #checked>
                          显示
                        </template>
                        <template #unchecked>
                          不显示
                        </template>
                      </n-switch>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="排序" path="sort">
                    <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" style="width: 100px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="状态" path="status" :show-require-mark="true">
                    <a-radio-group v-model:value="formValue.status" name="status">
                      <a-radio v-for="model in options.sys_normal_disable"
                               :key="model.value"
                               :value="model.value">{{model.label}}</a-radio>
                    </a-radio-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>

      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/pmsIndexNav';
import { options, State, newState } from './nav_model';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";

const emit = defineEmits(['reloadTable']);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});
const tagLanguage = ref({
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
      nameLanguage.value.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      nameLanguage.value.ko = ''
    }
  }else if(type == 'tag'){
    if(lang == 'zh'){
      // 简体中文
      tagLanguage.value.zh_CN = ''
    }else if(lang == 'en'){
      // 韩语
      tagLanguage.value.ko = ''
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
      }else if(type == 'tag'){
        if(lang == 'zh'){
          // 简体中文
          tagLanguage.value.zh_CN = _res.zh_CN
        }else if(lang == 'en'){
          // 韩语
          tagLanguage.value.ko = _res.ko
        }
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
    nameLanguage.value = {
      en: '',
      zh: '',
      ja: '',
      ko: '',
      zh_CN: '',
    }
    tagLanguage.value = {
      en: '',
      zh: '',
      ja: '',
      ko: '',
      zh_CN: '',
    }

    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;

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

      if(res.tagLanguage){
        const tagLanguageObj = jsontoobj(res.tagLanguage);
        tagLanguage.value.zh = tagLanguageObj.zh && tagLanguageObj.zh.content ? tagLanguageObj.zh.content : ''
        tagLanguage.value.en = tagLanguageObj.en && tagLanguageObj.en.content ? tagLanguageObj.en.content : ''
        tagLanguage.value.ko = tagLanguageObj.ko && tagLanguageObj.ko.content ? tagLanguageObj.ko.content : ''
        tagLanguage.value.ja = tagLanguageObj.ja && tagLanguageObj.ja.content ? tagLanguageObj.ja.content : ''
        if(tagLanguageObj.zh_CN){
          tagLanguage.value.zh_CN = tagLanguageObj.zh_CN.content ? tagLanguageObj.zh_CN.content : ''
        }else if(tagLanguageObj.zh_cn){
          tagLanguage.value.zh_CN = tagLanguageObj.zh_cn.content ? tagLanguageObj.zh_cn.content : ''
        }else{
          tagLanguage.value.zh_CN = ''
        }
      }
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  // 名称-多语言
  formValue.value.nameLanguage = nameLanguage.value

  // 名称-多语言
  formValue.value.tagLanguage = tagLanguage.value

  formValue.value.name = ""
  formValue.value.tag = ""

  if(!nameLanguage.value.zh || !nameLanguage.value.en || !nameLanguage.value.ja){
    formBtnLoading.value = false;
    message.error('名称请填写完整');
    return false;
  }

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
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

function closeForm() {
  showModal.value = false;
  loading.value = false;
  nameLanguage.value = {
    en: '',
    zh: '',
    ja: '',
    ko: '',
    zh_CN: '',
  }
  tagLanguage.value = {
    en: '',
    zh: '',
    ja: '',
    ko: '',
    zh_CN: '',
  }
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


