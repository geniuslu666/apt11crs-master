<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules" ref="formRef">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              分享图标
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="分享图标" path="icon">
                <FileChooser1 v-model:value="formValue.icon" :maxNumber="1"
                              fileType="default"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <n-divider title-placement="left">
              分享链接
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="分享链接" path="shareUrl">
                <n-input v-model:value="formValue.shareUrl" placeholder="分享链接"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="分享主域名" path="shareDomain">
                <n-input v-model:value="formValue.shareDomain" placeholder="分享主域名"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <n-divider title-placement="left">
              分享标题
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="分享标题-中文" path="titleZh">
                <n-input v-model:value="formValue.titleZh" placeholder="标题中文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="语言同步" path="language">
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
                  <img src="../../assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                  一键翻译
                </n-button>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享标题-繁体中文" path="titleZhCn">
                <n-input v-model:value="formValue.titleZhCn" placeholder="标题繁体"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享标题-日文" path="titleJa">
                <n-input v-model:value="formValue.titleJa" placeholder="标题日文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享标题-英文" path="titleEn">
                <n-input v-model:value="formValue.titleEn" placeholder="标题英文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享标题-韩文" path="titleKo">
                <n-input v-model:value="formValue.titleKo" placeholder="标题韩文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <n-divider title-placement="left">
              分享内容
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="分享内容-中文" path="titleZh">
                <n-input v-model:value="formValue.contentZh" placeholder="内容中文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="语言同步" path="language">
                <n-checkbox-group v-model:value="transContentItem" @update:value="handleTransContentItem">
                  <n-space item-style="display: flex;">
                    <n-checkbox value="ja" label="日语"/>
                    <n-checkbox value="en" label="英语"/>
                    <n-checkbox value="ko" label="韩语"/>
                    <n-checkbox value="zh_CN" label="繁体中文"/>
                  </n-space>
                </n-checkbox-group>
                <n-button :loading="transContBtnLoading" @click="handleContentTrans" type="primary" ghost
                          size="small">
                  <img src="../../assets/images/icon_translate.png" width="16" style="margin-right: 3px">
                  一键翻译
                </n-button>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享内容-繁体中文" path="titleZhCn">
                <n-input v-model:value="formValue.contentZhCn" placeholder="内容繁体"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享内容-日文" path="titleJa">
                <n-input v-model:value="formValue.contentJa" placeholder="内容日文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享内容-英文" path="titleEn">
                <n-input v-model:value="formValue.contentEn" placeholder="内容英文"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="分享内容-韩文" path="titleKo">
                <n-input v-model:value="formValue.contentKo" placeholder="内容韩文"/>
              </n-form-item>
            </div>
          </n-gi>
        </n-grid>
        <div style="text-align: center">
          <n-space justify="center">
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';
import FileChooser from "@/components/FileChooser/index.vue";
import {Text} from "@/api/translate";

const formBtnLoading = ref(false);
const group = ref('wxShare');
const show = ref(false);
const formRef: any = ref(null);
const message = useMessage();
const dialog = useDialog();
const transItem = ref(['en', 'ja', 'ko', 'zh', 'zh_CN']);
const transContentItem = ref(['en', 'ja', 'ko', 'zh', 'zh_CN']);
const transBtnLoading = ref(false);
const transContBtnLoading = ref(false);

const formValue = ref({
  titleZh: "",
  titleZhCn: "",
  titleJa: "",
  titleEn: "",
  titleKo: "",
  contentZh: "",
  contentZhCn: "",
  contentJa: "",
  contentEn: "",
  contentKo: "",
  icon: "",
  shareUrl: "",
  shareDomain: "",
});

const rules = {
  /*  basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },*/
};

function formSubmit() {
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      updateConfig({group: group.value, list: formValue.value}).then((_res) => {
        formBtnLoading.value = false;
        message.success('更新成功');
        load();
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('验证失败，请填写完整信息');
    }
  });
}

onMounted(() => {
  load();
});

function handleTransItem(value) {
  console.log(value)
  transItem.value = value
}

function handleTransContentItem(value) {
  console.log(value)
  transContentItem.value = value
}

function load() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    getConfig({group: group.value})
      .then((res) => {
        // payMode.value = res.list.paymode.split(",").map(item => String(item));
        formValue.value = res.list;
      })
      .finally(() => {
        show.value = false;
      });
  });
}

function handleTrans() {
  transBtnLoading.value = true;
  if (!formValue.value.titleZh) {
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
        text: formValue.value.titleZh
      }).then((_res) => {
        message.success('翻译成功');
        if (transItem.value.indexOf('ja') !== -1) {
          formValue.value.titleJa = _res.ja
        }
        if (transItem.value.indexOf('ko') !== -1) {
          formValue.value.titleKo = _res.ko
        }
        if (transItem.value.indexOf('en') !== -1) {
          formValue.value.titleEn = _res.en
        }
        if (transItem.value.indexOf('zh_CN') !== -1) {
          formValue.value.titleZhCn = _res.zh_CN
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

function handleContentTrans() {
  transContBtnLoading.value = true;
  if (!formValue.value.contentZh) {
    message.error('请先填写简体中文的内容');
    transContBtnLoading.value = false;
    return false;
  }
  if (transItem.value.length <= 0) {
    message.error('请选择要进行翻译的语言');
    transContBtnLoading.value = false;
    return false;
  }

  dialog.warning({
    title: '警告',
    content: '翻译会覆盖已选语言下所填写的内容，你确定要进行翻译吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Text({
        text: formValue.value.contentZh
      }).then((_res) => {
        message.success('翻译成功');
        if (transContentItem.value.indexOf('ja') !== -1) {
          formValue.value.contentJa = _res.ja
        }
        if (transContentItem.value.indexOf('ko') !== -1) {
          formValue.value.contentKo = _res.ko
        }
        if (transContentItem.value.indexOf('en') !== -1) {
          formValue.value.contentEn = _res.en
        }
        if (transContentItem.value.indexOf('zh_CN') !== -1) {
          formValue.value.contentZhCn = _res.zh_CN
        }
        setTimeout(() => {
          transContBtnLoading.value = false;
        });
      }).catch((err) => {
        message.error('翻译失败');
        transContBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      transContBtnLoading.value = false;
    }
  });

}

</script>
