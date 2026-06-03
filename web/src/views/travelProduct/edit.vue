<template>
  <div>
    <n-drawer v-model:show="showModal" :trap-focus="false" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑产品 #' + formValue.id : '新增产品' }}</div>
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
            require-mark-placement="right-hanging"
          >
            <n-tabs type="line" animated v-model:value="tabValue">
              <n-tab-pane name="base" tab="基础设置">
                <div class="level-detail-div">
                  <div class="level-detail-div-title">
                    <div></div>
                    基本信息
                  </div>
                </div>
                <n-grid cols="1 600:3" x-gap="80">
                  <n-gi>
                    <n-form-item label="产品标题_简体中文" path="name_zh">
                      <n-input placeholder="简体中文产品标题" v-model:value="titleLanguage.zh" :style="{ width: '300px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="产品标题_日本语" path="name_ja">
                      <n-input placeholder="日本语产品标题" v-model:value="titleLanguage.ja" :style="{ width: '300px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="产品标题_English" path="name_en">
                      <n-input placeholder="英语产品标题" v-model:value="titleLanguage.en" :style="{ width: '300px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="产品标题_한국어" path="name_ko">
                      <n-input placeholder="韩语产品标题" v-model:value="titleLanguage.ko" :style="{ width: '300px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi span="1 600:2">
                    <n-form-item label="产品标题_繁体中文" path="name_zh_CN">
                      <n-input placeholder="繁体中文产品标题" v-model:value="titleLanguage.zh_CN" :style="{ width: '300px' }" />
                    </n-form-item>
                  </n-gi>

                  <n-gi>
                    <n-form-item label="副标题_简体中文" path="subname_zh">
                      <n-input placeholder="简体中文副标题" v-model:value="subTitleLanguage.zh" :style="{ width: '400px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="副标题_日本语" path="subname_ja">
                      <n-input placeholder="日本语副标题" v-model:value="subTitleLanguage.ja" :style="{ width: '400px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="副标题_English" path="subname_en">
                      <n-input placeholder="英语副标题" v-model:value="subTitleLanguage.en" :style="{ width: '400px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="副标题_한국어" path="subname_ko">
                      <n-input placeholder="韩语副标题" v-model:value="subTitleLanguage.ko" :style="{ width: '400px' }" />
                    </n-form-item>
                  </n-gi>
                  <n-gi span="1 600:2">
                    <n-form-item label="副标题_繁体中文" path="subname_zh_CN">
                      <n-input placeholder="繁体中文副标题" v-model:value="subTitleLanguage.zh_CN" :style="{ width: '400px' }" />
                    </n-form-item>
                  </n-gi>

                  <n-gi span="1 600:3">
                    <n-form-item label="列表图" path="listImage">
                      <FileChooser1 v-model:value="formValue.listImage" :maxNumber="1" fileType="default" />
                    </n-form-item>
                  </n-gi>

                  <n-gi span="1 600:3">
                    <n-form-item label="图集" path="carouselImages">
                      <FileChooser :maxNumber="10" v-model:value="formValue.carouselImages" />
                    </n-form-item>
                    <!--                <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 150px;">建议尺寸：638px*300px</div>-->
                  </n-gi>
<!--                  <n-gi span="1 600:3">
                    <n-form-item label="售价" path="price">
                      <n-input-group>
                        <n-input-number placeholder="请输入售价" :min="0" v-model:value="formValue.price" style="width: 120px" />
                        <n-input-group-label>JPY</n-input-group-label>
                      </n-input-group>
                    </n-form-item>
                  </n-gi>
                  <n-gi span="1 600:3">
                    <n-form-item label="每日接待人数" path="dailyCapacity">
                      <n-input-group>
                        <n-input-number placeholder="请输入每日最大接待人数" :min="0" v-model:value="formValue.dailyCapacity" style="width: 120px" />
                        <n-input-group-label>人</n-input-group-label>
                      </n-input-group>
                    </n-form-item>
                  </n-gi>-->
                  <n-gi span="1 600:3">
                    <n-form-item label="总库存" path="dailyCapacity">
                      <n-input-group>
                        <n-input-number placeholder="请输入总库存" :min="0" v-model:value="formValue.stock" style="width: 120px" />
                        <n-input-group-label>人</n-input-group-label>
                      </n-input-group>
                    </n-form-item>
                  </n-gi>
                  <n-gi span="1 600:3">
                    <n-form-item label="状态" path="status">
                      <n-radio-group v-model:value="formValue.status" name="status">
                        <n-radio-button
                          :value="1"
                          label="启用"
                        />
                        <n-radio-button
                          :value="2"
                          label="禁用"
                        />
                      </n-radio-group>
                    </n-form-item>
                  </n-gi>
                  <n-gi span="1 600:3">
                    <n-form-item label="排序">
                      <n-input-number v-model:value="formValue.sort" :min="0" placeholder="越大越靠前" style="width: 100px" />
                    </n-form-item>
                  </n-gi>
                </n-grid>

                <div class="level-detail-div">
                  <div class="level-detail-div-title">
                    <div></div>
                    预约设置
                  </div>
                </div>
                <n-grid cols="1" x-gap="80">
                  <n-gi>
                    <n-form-item label="最大可预约天数">
                      <n-input-group>
                        <n-input-group-label>最长可预约</n-input-group-label>
                        <n-input-number placeholder="请输入天数" :min="0" :precision="0" :show-button="false" v-model:value="formValue.maxBookDays" style="width: 100px" />
                        <n-input-group-label>天后</n-input-group-label>
                      </n-input-group>
                      <!--                  <n-input-number v-model:value="formValue.maxBookDays" :min="1" placeholder="请输入最大可预约天数" style="width: 100%" />-->
                    </n-form-item>
                  </n-gi>
                  <n-gi>
                    <n-form-item label="至少提前预订天数">
                      <n-input-group>
                        <n-input-group-label>需至少提前</n-input-group-label>
                        <n-input-number placeholder="请输入天数" :min="1" :precision="0" v-model:value="formValue.advanceBookDays" style="width: 100px" />
                        <n-input-group-label>天</n-input-group-label>
                      </n-input-group>
                      <!--                  <n-input-number v-model:value="formValue.advanceBookDays" :min="0" placeholder="请输入至少提前预订天数" style="width: 100%" />-->
                    </n-form-item>
                  </n-gi>
                </n-grid>
              </n-tab-pane>
              <n-tab-pane name="contentTab" tab="内容介绍">
                <n-tabs v-model:value="contentLangTab" type="card" style="margin-bottom: 16px">
                  <n-tab-pane name="zh" tab="简体中文">
                    <n-form-item path="contentZh">
                      <Editor style="height: 300px" id="contentZh" v-model:modelValue="formValue.contentZh" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="en" tab="English">
                    <n-form-item path="contentEn">
                      <Editor style="height: 300px" id="contentEn" v-model:modelValue="formValue.contentEn" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="ja" tab="日本語">
                    <n-form-item path="contentJa">
                      <Editor style="height: 300px" id="contentJa" v-model:modelValue="formValue.contentJa" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="ko" tab="한국어">
                    <n-form-item path="contentKo">
                      <Editor style="height: 300px" id="contentKo" v-model:modelValue="formValue.contentKo" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="tw" tab="繁体中文">
                    <n-form-item path="contentTw">
                      <Editor style="height: 300px" id="contentTw" v-model:modelValue="formValue.contentTw" />
                    </n-form-item>
                  </n-tab-pane>
                </n-tabs>
              </n-tab-pane>
              <n-tab-pane name="tripPlanningTab" tab="行程规划">
                <n-tabs v-model:value="tripPlanningLangTab" type="card" style="margin-bottom: 16px">
                  <n-tab-pane name="tripPlanningZh" tab="简体中文">
                    <n-form-item path="tripPlanningZh">
                      <Editor style="height: 300px" id="tripPlanningZh" v-model:modelValue="formValue.tripPlanningZh" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="tripPlanningEn" tab="English">
                    <n-form-item path="tripPlanningEn">
                      <Editor style="height: 300px" id="tripPlanningEn" v-model:modelValue="formValue.tripPlanningEn" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="tripPlanningJa" tab="日本語">
                    <n-form-item path="tripPlanningJa">
                      <Editor style="height: 300px" id="tripPlanningJa" v-model:modelValue="formValue.tripPlanningJa" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="tripPlanningKo" tab="한국어">
                    <n-form-item path="tripPlanningKo">
                      <Editor style="height: 300px" id="tripPlanningKo" v-model:modelValue="formValue.tripPlanningKo" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="tripPlanningTw" tab="繁体中文">
                    <n-form-item path="tripPlanningTw">
                      <Editor style="height: 300px" id="tripPlanningTw" v-model:modelValue="formValue.tripPlanningTw" />
                    </n-form-item>
                  </n-tab-pane>
                </n-tabs>
              </n-tab-pane>
              <n-tab-pane name="bookingNotesTab" tab="预约须知">
                <n-tabs v-model:value="bookingNotesLangTab" type="card" style="margin-bottom: 16px">
                  <n-tab-pane name="bookingNotesZh" tab="简体中文">
                    <n-form-item path="bookingNotesZh">
                      <Editor style="height: 300px" id="bookingNotesZh" v-model:modelValue="formValue.bookingNotesZh" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="bookingNotesEn" tab="English">
                    <n-form-item path="bookingNotesEn">
                      <Editor style="height: 300px" id="bookingNotesEn" v-model:modelValue="formValue.bookingNotesEn" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="bookingNotesJa" tab="日本語">
                    <n-form-item path="bookingNotesJa">
                      <Editor style="height: 300px" id="bookingNotesJa" v-model:modelValue="formValue.bookingNotesJa" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="bookingNotesKo" tab="한국어">
                    <n-form-item path="bookingNotesKo">
                      <Editor style="height: 300px" id="bookingNotesKo" v-model:modelValue="formValue.bookingNotesKo" />
                    </n-form-item>
                  </n-tab-pane>
                  <n-tab-pane name="bookingNotesTw" tab="繁体中文">
                    <n-form-item path="bookingNotesTw">
                      <Editor style="height: 300px" id="bookingNotesTw" v-model:modelValue="formValue.bookingNotesTw" />
                    </n-form-item>
                  </n-tab-pane>
                </n-tabs>
              </n-tab-pane>
              <!-- 车型管理已移至产品列表页的"车型管理"按钮 -->
            </n-tabs>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed} from 'vue';
import { useMessage } from 'naive-ui';
import { View, Edit } from '@/api/travelProduct';
import { State, newState, rules } from './model';
import {adaModalWidth} from "@/utils/hotgo";
import {jsontoobj} from "@/utils/smjcomm";
import Editor from "@/components/Editor/editor.vue";

const emit = defineEmits(['reloadTable']);
const tabValue = ref('base')
const message = useMessage();
const formRef = ref<any>({});
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const showModal = ref(false);
const formBtnLoading = ref(false);
const formValue = ref<State>(newState(null));

const titleLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const subTitleLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const contentLangTab = ref('zh');
const tripPlanningLangTab = ref('tripPlanningZh')
const bookingNotesLangTab = ref('bookingNotesZh')

function confirmForm(e) {
  formValue.value.titleLanguage = titleLanguage.value
  formValue.value.subTitleLanguage = subTitleLanguage.value

  formValue.value.title = ""
  formValue.value.subTitle = ""

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value?.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch(() => {
        message.error('操作失败，请稍后再试');
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

async function getInfo(id) {
  const res = await View({ id: id, isLanguage: true });
  formValue.value = res;

  formValue.value.carouselImages = res.carouselImages ? JSON.parse(res.carouselImages) : []

  if (res.titleLanguage) {
    res.titleLanguage = jsontoobj(res.titleLanguage);
    titleLanguage.value.zh = res.titleLanguage.zh && res.titleLanguage.zh.content ? res.titleLanguage.zh.content : ''
    titleLanguage.value.en = res.titleLanguage.en && res.titleLanguage.en.content ? res.titleLanguage.en.content : ''
    titleLanguage.value.ko = res.titleLanguage.ko && res.titleLanguage.ko.content ? res.titleLanguage.ko.content : ''
    titleLanguage.value.ja = res.titleLanguage.ja && res.titleLanguage.ja.content ? res.titleLanguage.ja.content : ''
    if (res.titleLanguage.zh_CN) {
      titleLanguage.value.zh_CN = res.titleLanguage.zh_CN.content ? res.titleLanguage.zh_CN.content : ''
    } else if (res.titleLanguage.zh_cn) {
      titleLanguage.value.zh_CN = res.titleLanguage.zh_cn.content ? res.titleLanguage.zh_cn.content : ''
    } else {
      titleLanguage.value.zh_CN = ''
    }
  }

  if (res.subTitleLanguage) {
    res.subTitleLanguage = jsontoobj(res.subTitleLanguage);
    subTitleLanguage.value.zh = res.subTitleLanguage.zh && res.subTitleLanguage.zh.content ? res.subTitleLanguage.zh.content : ''
    subTitleLanguage.value.en = res.subTitleLanguage.en && res.subTitleLanguage.en.content ? res.subTitleLanguage.en.content : ''
    subTitleLanguage.value.ko = res.subTitleLanguage.ko && res.subTitleLanguage.ko.content ? res.subTitleLanguage.ko.content : ''
    subTitleLanguage.value.ja = res.subTitleLanguage.ja && res.subTitleLanguage.ja.content ? res.subTitleLanguage.ja.content : ''
    if (res.subTitleLanguage.zh_CN) {
      subTitleLanguage.value.zh_CN = res.subTitleLanguage.zh_CN.content ? res.subTitleLanguage.zh_CN.content : ''
    } else if (res.titleLanguage.zh_cn) {
      subTitleLanguage.value.zh_CN = res.subTitleLanguage.zh_cn.content ? res.subTitleLanguage.zh_cn.content : ''
    } else {
      subTitleLanguage.value.zh_CN = ''
    }
  }
}


async function openModal(state: State) {
  contentLangTab.value = 'zh';
  tripPlanningLangTab.value = 'tripPlanningZh';
  bookingNotesLangTab.value = 'bookingNotesZh';
  showModal.value = true;
  loading.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    titleLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };
    subTitleLanguage.value = {
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
  loading.value = false;

}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
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
