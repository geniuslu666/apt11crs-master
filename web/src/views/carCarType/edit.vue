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
        <n-card
          :bordered="false"
          class="proCard mt-4"
          size="small"
          :segmented="{ content: true }"
          :title="formValue.id > 0 ? '编辑车型 #' + formValue.id : '添加车型'"
        >
          <n-tabs type="line" animated v-model:value="tabValue">
            <n-tab-pane name="base" tab="基础信息">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="车型名称_简体中文" path="name_zh" :show-require-mark="true">
                    <n-input placeholder="请输入车型名称" v-model:value="nameLanguage.zh" @blur="translate('name','zh',nameLanguage.zh)" style="width: 250px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型内容_日本语" path="name_ja">
                    <n-input placeholder="请输入车型名称" v-model:value="nameLanguage.ja" style="width: 250px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型内容_English" path="name_en">
                    <n-input placeholder="请输入车型名称" v-model:value="nameLanguage.en" style="width: 250px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型简介_简体中文" path="content_zh" :show-require-mark="true">
                    <n-input placeholder="请输入车型简介" v-model:value="descLanguage.zh" @blur="translate('desc','zh',descLanguage.zh)" style="width: 400px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型内容_日本语" path="content_ja">
                    <n-input placeholder="请输入车型简介" v-model:value="descLanguage.ja" style="width: 400px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型内容_English" path="content_en">
                    <n-input placeholder="请输入车型简介" v-model:value="descLanguage.en" style="width: 400px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="车型图片" path="photo" :show-require-mark="true" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.image" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：180px*120px</div>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="座位数" path="seatNum">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.seatNum" style="width: 100px" />
                      <n-input-group-label>座</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="建议乘员人数" path="passengerNum">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.passengerNum" style="width: 100px" />
                      <n-input-group-label>人</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="最大容纳行李件数" path="maxPackageNum">
                    <n-input-group>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.maxPackageNum" style="width: 100px" />
                      <n-input-group-label>件</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="1">
                  <n-form-item label="是否支持儿童座椅" path="childrenSeatSupport">
                    <n-radio-group v-model:value="formValue.childrenSeatSupport" name="childrenSeatSupport">
                      <n-space>
                        <n-radio value="Y">
                          支持
                        </n-radio>
                        <n-radio value="N">
                          不支持
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item label=" " path="childrenSeatSpace" v-if="formValue.childrenSeatSupport == 'Y'">
                    <n-input-group>
                      <n-input-group-label>1个儿童座椅占用</n-input-group-label>
                      <n-input-number placeholder="请输入" :show-button="false" :min="0" v-model:value="formValue.childrenSeatSpace" style="width: 100px" />
                      <n-input-group-label>个座位</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>
                <!--              <n-gi span="1">-->
                <!--                <n-form-item label="状态" path="status">-->
                <!--                  <n-radio-group v-model:value="formValue.status" name="status">-->
                <!--                    <n-radio-button-->
                <!--                      v-for="status in options.sys_normal_disable"-->
                <!--                      :key="status.value"-->
                <!--                      :value="status.value"-->
                <!--                      :label="status.label"-->
                <!--                    />-->
                <!--                  </n-radio-group>-->
                <!--                </n-form-item>-->
                <!--              </n-gi>-->
                <n-gi span="1">
                  <n-form-item label="排序" path="sort">
                    <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </n-tab-pane>

            <n-tab-pane name="content" tab="车型详情">
              <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
                <n-gi span="1">
                  <n-form-item label="内容" path="content">
                    <Editor id="content" v-model:modelValue="formValue.content" style="height: 450px" />
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
import {onMounted, reactive, ref} from 'vue';
import { Edit, View, MaxSort } from '@/api/carCarType';
import { options, State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {Text} from "@/api/translate";
import {jsontoobj} from "@/utils/smjcomm";
import UploadImage from "@/components/Upload/uploadImage.vue";
import {useRouter} from "vue-router";
import {useTabsViewStore} from "@/store/modules/tabsView";
import Editor from "@/components/Editor/editor.vue";

const message = useMessage();
const settingStore = useProjectSettingStore();
const tabValue = ref('base')
const show = ref(false);
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const params = router.currentRoute.value.params;
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const nameLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});
const descLanguage = reactive({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

function translate(type,lang,text){
  if(lang == 'zh'){
    if(type == 'name'){
      // 简体中文
      nameLanguage.zh_CN = ''
    }else{
      // 简体中文
      descLanguage.zh_CN = ''
    }
  }
  if(text){
    Text({
      text: text
    }).then((_res) => {
      if(lang == 'zh'){
        if(type == 'name'){
          // 简体中文
          nameLanguage.zh_CN = _res.zh_CN
        }else{
          // 简体中文
          descLanguage.zh_CN = _res.zh_CN
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
      nameLanguage.ko = nameLanguage.en
      descLanguage.ko = descLanguage.en
      formValue.value.nameLanguage = nameLanguage
      formValue.value.descLanguage = descLanguage
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'carCarTypeIndex', params: {  } });
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

  if(detailRes.descLanguage){
    const descLanguageObj = jsontoobj(detailRes.descLanguage);
    descLanguage.zh = descLanguageObj.zh && descLanguageObj.zh.content ? descLanguageObj.zh.content : ''
    descLanguage.en = descLanguageObj.en && descLanguageObj.en.content ? descLanguageObj.en.content : ''
    descLanguage.ko = descLanguageObj.ko && descLanguageObj.ko.content ? descLanguageObj.ko.content : ''
    descLanguage.ja = descLanguageObj.ja && descLanguageObj.ja.content ? descLanguageObj.ja.content : ''
    if(descLanguageObj.zh_CN){
      descLanguage.zh_CN = descLanguageObj.zh_CN.content ? descLanguageObj.zh_CN.content : ''
    }else if(descLanguageObj.zh_cn){
      descLanguage.zh_CN = descLanguageObj.zh_cn.content ? descLanguageObj.zh_cn.content : ''
    }else{
      descLanguage.zh_CN = ''
    }
  }
}

async function loadMax(){
  let max = await MaxSort()
  formValue.value.sort = max.sort
}

async function init(){
  formValue.value = newState(null);
  nameLanguage.zh = '';
  nameLanguage.en = '';
  nameLanguage.ja = '';
  nameLanguage.ko = '';
  nameLanguage.zh_CN = '';

  descLanguage.zh = '';
  descLanguage.en = '';
  descLanguage.ja = '';
  descLanguage.ko = '';
  descLanguage.zh_CN = '';
}

onMounted(async() => {
  show.value = true;
  if(parseInt(params.id) > 0){
    await load();
  }else{
    await init();
    await loadMax()
  }

  show.value = false;
})
</script>

<style lang="less"></style>


