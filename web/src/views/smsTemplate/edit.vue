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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑通知模版 #' + formValue.id : '添加通知模版' }}</div>
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
                基础信息
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="场景" path="scene">
                    <n-radio-group v-model:value="formValue.scene">
                      <n-space>
                        <n-radio :value="1">
                          会员
                        </n-radio>
                        <n-radio :value="2">
                          住宿
                        </n-radio>
                        <n-radio :value="3">
                          接送机
                        </n-radio>
                        <n-radio :value="4">
                          按摩
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="别名" path="event">
                    <n-input placeholder="请输入别名" v-model:value="formValue.event" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="标题" path="title">
                    <n-input placeholder="请输入标题" v-model:value="formValue.title" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="内容_简体中文" path="name_zh" >
                    <n-input
                      type="textarea"
                      placeholder="简体中文内容"
                      autosize style="min-height: 100px"
                      v-model:value="contentLanguage.zh"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="内容_日本语" path="name_ja" >
                    <n-input
                      type="textarea"
                      placeholder="日本语内容"
                      autosize style="min-height: 100px"
                      v-model:value="contentLanguage.ja"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="内容_English" path="name_en" >
                    <n-input
                      type="textarea"
                      placeholder="英语内容"
                      autosize style="min-height: 100px"
                      v-model:value="contentLanguage.en"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="内容_한국어" path="name_ko" >
                    <n-input
                      type="textarea"
                      placeholder="韩语内容"
                      autosize style="min-height: 100px"
                      v-model:value="contentLanguage.ko"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="内容_繁体中文" path="name_zh_CN" >
                    <n-input
                      type="textarea"
                      placeholder="繁体中文内容"
                      autosize style="min-height: 100px"
                      v-model:value="contentLanguage.zh_CN"
                    />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>

            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                一信通
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="模版ID" path="umsId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.umsId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="模版内容" path="umsContent" >
                    <n-input
                      type="textarea"
                      placeholder="请输入模版内容"
                      autosize style="min-height: 100px"
                      v-model:value="formValue.umsContent"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="变量" path="umsParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.umsParam" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>

            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                腾讯云
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="日文模版ID" path="tencentJaId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.tencentJaId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="日文模版内容" path="tencentJaContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.tencentJaContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="日文变量" path="tencentJaParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.tencentJaParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>
                <n-gi style="margin-top: 24px">
                  <n-form-item label="英文模版ID" path="tencentEnId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.tencentEnId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="英文模版内容" path="tencentEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.tencentEnContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="英文变量" path="tencentEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.tencentEnParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>

                <n-gi style="margin-top: 24px">
                  <n-form-item label="韩文模版ID" path="tencentKoId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.tencentKoId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="韩文模版内容" path="tencentEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.tencentKoContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="韩文变量" path="tencentEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.tencentKoParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>

                <n-gi style="margin-top: 24px">
                  <n-form-item label="繁体模版ID" path="tencentTwId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.tencentTwId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="繁体模版内容" path="tencentEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.tencentTwContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="繁体变量" path="tencentEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.tencentTwParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>

            <div class="level-detail-div" style="margin-top: 24px">
              <div class="level-detail-div-title">
                <div></div>
                阿里云
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="日文模版ID" path="aliyunJaId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.aliyunJaId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="日文模版内容" path="aliyunJaContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.aliyunJaContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="日文变量" path="aliyunJaParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.aliyunJaParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>
                <n-gi style="margin-top: 24px">
                  <n-form-item label="英文模版ID" path="aliyunEnId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.aliyunEnId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="英文模版内容" path="aliyunEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.aliyunEnContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="英文变量" path="aliyunEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.aliyunEnParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>

                <n-gi style="margin-top: 24px">
                  <n-form-item label="韩文模版ID" path="aliyunKoId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.aliyunKoId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="韩文模版内容" path="aliyunEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.aliyunKoContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="韩文变量" path="aliyunEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.aliyunKoParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
                  </n-form-item>
                </n-gi>

                <n-gi style="margin-top: 24px">
                  <n-form-item label="繁体模版ID" path="aliyunTwId">
                    <n-input placeholder="请输入模版ID" v-model:value="formValue.aliyunTwId" :style="{ width: '300px' }"/>
                  </n-form-item>
                </n-gi>
<!--                <n-gi>-->
<!--                  <n-form-item label="繁体模版内容" path="aliyunEnContent" >-->
<!--                    <n-input placeholder="请输入模版内容" v-model:value="formValue.aliyunTwContent" />-->
<!--                  </n-form-item>-->
<!--                </n-gi>-->
                <n-gi>
                  <n-form-item label="繁体变量" path="aliyunEnParam" >
                    <n-input placeholder="请输入变量" v-model:value="formValue.aliyunTwParam" />
                    <template #feedback>
                      <div style="font-size: 12px; color: red">按顺序填写，多个用,分隔</div>
                    </template>
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
import { Edit, View } from '@/api/smsTemplate';
import {  State, newState } from './model';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import {jsontoobj} from "@/utils/smjcomm";

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

const contentLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    contentLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
      if(res.contentLanguage){
        res.contentLanguage = jsontoobj(res.contentLanguage);
        contentLanguage.value.zh = res.contentLanguage.zh && res.contentLanguage.zh.content ? res.contentLanguage.zh.content : ''
        contentLanguage.value.en = res.contentLanguage.en && res.contentLanguage.en.content ? res.contentLanguage.en.content : ''
        contentLanguage.value.ko = res.contentLanguage.ko && res.contentLanguage.ko.content ? res.contentLanguage.ko.content : ''
        contentLanguage.value.ja = res.contentLanguage.ja && res.contentLanguage.ja.content ? res.contentLanguage.ja.content : ''
        if(res.contentLanguage.zh_CN){
          contentLanguage.value.zh_CN = res.contentLanguage.zh_CN.content ? res.contentLanguage.zh_CN.content : ''
        }else if(res.contentLanguage.zh_cn){
          contentLanguage.value.zh_CN = res.contentLanguage.zh_cn.content ? res.contentLanguage.zh_cn.content : ''
        }else{
          contentLanguage.value.zh_CN = ''
        }
      }

      formValue.value.umsId = res.umsTemplate?.id
      formValue.value.umsContent = res.umsTemplate?.content
      formValue.value.umsParam = res.umsTemplate?.param

      formValue.value.tencentJaId = res.tencentTemplate?.ja?.id
      formValue.value.tencentJaParam = res.tencentTemplate?.ja?.param
      formValue.value.tencentKoId = res.tencentTemplate?.ko?.id
      formValue.value.tencentKoParam = res.tencentTemplate?.ko?.param
      formValue.value.tencentEnId = res.tencentTemplate?.en?.id
      formValue.value.tencentEnParam = res.tencentTemplate?.en?.param
      formValue.value.tencentTwId = res.tencentTemplate?.tw?.id
      formValue.value.tencentTwParam = res.tencentTemplate?.tw?.param

      formValue.value.aliyunJaId = res.aliyunTemplate?.ja?.id
      formValue.value.aliyunJaParam = res.aliyunTemplate?.ja?.param
      formValue.value.aliyunKoId = res.aliyunTemplate?.ko?.id
      formValue.value.aliyunKoParam = res.aliyunTemplate?.ko?.param
      formValue.value.aliyunEnId = res.aliyunTemplate?.en?.id
      formValue.value.aliyunEnParam = res.aliyunTemplate?.en?.param
      formValue.value.aliyunTwId = res.aliyunTemplate?.tw?.id
      formValue.value.aliyunTwParam = res.aliyunTemplate?.tw?.param
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
      // 内容-多语言
      formValue.value.contentLanguage = contentLanguage.value

      // 一信通JSON处理
      formValue.value.umsTemplate = JSON.stringify({
        id: formValue.value.umsId,
        content: formValue.value.umsContent,
        param: formValue.value.umsParam
      })

      // 腾讯云JSON处理
      formValue.value.tencentTemplate = JSON.stringify({
        ja: {
          id: formValue.value.tencentJaId,
          param: formValue.value.tencentJaParam
        },
        ko: {
          id: formValue.value.tencentKoId,
          param: formValue.value.tencentKoParam
        },
        en: {
          id: formValue.value.tencentEnId,
          param: formValue.value.tencentEnParam
        },
        tw: {
          id: formValue.value.tencentTwId,
          param: formValue.value.tencentTwParam
        }
      })

      // 阿里云JSON处理
      formValue.value.aliyunTemplate = JSON.stringify({
        ja: {
          id: formValue.value.aliyunJaId,
          param: formValue.value.aliyunJaParam
        },
        ko: {
          id: formValue.value.aliyunKoId,
          param: formValue.value.aliyunKoParam
        },
        en: {
          id: formValue.value.aliyunEnId,
          param: formValue.value.aliyunEnParam
        },
        tw: {
          id: formValue.value.aliyunTwId,
          param: formValue.value.aliyunTwParam
        }
      })


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


