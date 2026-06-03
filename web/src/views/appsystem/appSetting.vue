<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              APP版本信息
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="机型选择" path="phoneModelArr">
                <n-dynamic-tags v-model:value="phoneModelArr" :on-update:value="handleUpdateVersion"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-tabs v-model:value="tabValue" animated type="line">
                <n-tab-pane v-for="versionItem in appVersionArr" :key="versionItem.name"
                            :name="versionItem.name">
                  <n-form-item label="App版本号" path="appversion" style="margin-top: 18px">
                    <n-input v-model:value="versionItem.info.app_version"
                             placeholder="请输入App版本号"/>
                  </n-form-item>
                  <n-form-item label="App下载链接" path="app_privacy" style="margin-top: 6px" v-if="versionItem.name != 'miniapp' && versionItem.name != 'h5'">
                    <n-input v-model:value="versionItem.info.app_download"
                             placeholder="请输入App下载链接"/>
                  </n-form-item>
                  <n-form-item label="App新版本内容" path="app_version_content" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_version_content"
                             placeholder="请输入App版本内容" type="textarea"/>
                  </n-form-item>
                  <n-form-item label="App开启访问" path="app_open" style="margin-top: 6px" v-if="versionItem.name != 'miniapp' && versionItem.name != 'h5'">
                    <n-switch
                      v-model:value="versionItem.info.app_open"
                      size="large"
                      @update:value="versionOpenChange"
                    />
                  </n-form-item>
                  <n-form-item label="App审核" path="app_audit" style="margin-top: 6px" v-if="versionItem.name != 'miniapp' && versionItem.name != 'h5'">
                    <n-switch
                      v-model:value="versionItem.info.app_audit"
                      size="large"
                      @update:value="versionOpenChange3"
                    />
                  </n-form-item>
                  <n-form-item label="App强制升级" path="app_forced_upload" style="margin-top: 6px" v-if="versionItem.name != 'miniapp' && versionItem.name != 'h5'">
                    <n-switch
                      v-model:value="versionItem.info.app_forced_upload"
                      size="large"
                      @update:value="versionOpenChange2"
                    />
                  </n-form-item>
                  <n-form-item label="App支付方式" path="app_pay_type" style="margin-top: 6px" v-if="versionItem.name != 'miniapp'">
                    <n-checkbox-group v-model:value="versionItem.info.app_pay_model">
                      <n-space align="center" item-style="display: flex;">
                        <n-checkbox label="【PayCloud】支付宝" value="Alipay+" />
                        <n-checkbox label="【PayCloud】微信APP支付" value="WeChatPay" />
                        <n-checkbox label="Paypal" value="Paypal" />
                        <n-checkbox label="Paypal信用卡" value="PaypalCard" />
                        <n-checkbox label="Stripe信用卡" value="StripeCard" />
                        <n-checkbox label="【PayCloud】微信小程序" value="WeChatMiniPay" />
                        <n-checkbox label="【PayCloud】信用卡" value="CreditLinkPay" />
                        <n-checkbox label="PaypayH5" value="Paypay_h5" />
                        <n-checkbox label="PaypayApp" value="Paypay_app" />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                  <n-form-item label="支付方式" path="app_pay_type" style="margin-top: 6px" v-if="versionItem.name == 'miniapp'">
                    <n-radio-group v-model:value="versionItem.info.minipaychannel">
                      <n-space align="center" item-style="display: flex;">
                        <n-radio label="小程序支付" value="mlilife" />
                        <n-radio label="PayCloud支付" value="paycloud" />
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item label="App隐私协议-中文" path="app_privacy_zh" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_privacy_zh"
                             placeholder="请输入App隐私协议"/>
                  </n-form-item>

                  <n-form-item label="App隐私协议-繁体中文" path="app_privacy_zh_CN" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_privacy_zh_CN"
                             placeholder="请输入App隐私协议"/>
                  </n-form-item>
                  <n-form-item label="App隐私协议-英文" path="app_privacy_en" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_privacy_en"
                             placeholder="请输入App隐私协议"/>
                  </n-form-item>
                  <n-form-item label="App隐私协议-日文" path="app_privacy_ja" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_privacy_ja"
                             placeholder="请输入App隐私协议"/>
                  </n-form-item>
                  <n-form-item label="App隐私协议-韩文" path="app_privacy_ko" style="margin-top: 6px">
                    <n-input v-model:value="versionItem.info.app_privacy_ko"
                             placeholder="请输入App隐私协议"/>
                  </n-form-item>
                </n-tab-pane>
              </n-tabs>
            </div>
          </n-gi>
          <!-- <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="安卓下载链接" path="androidDownloadUrl">
                <n-input v-model:value="formValue.androidDownloadUrl" placeholder="安卓下载链接"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="苹果下载链接" path="iosDownloadUrl">
                <n-input v-model:value="formValue.iosDownloadUrl" placeholder="苹果下载链接"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="App新版本内容" path="appversionContent">
                <n-input
                  v-model:value="formValue.appversionContent"
                  :auto-size="{ minRows: 2, maxRows: 5 }"
                  placeholder="App新版本内容"
                  type="textarea"
                />
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="App开启访问" path="appOpen">
                <n-switch
                  v-model:value="formValue.appOpen"
                  size="large"
                  @update:value="systemOpenChange"
                />
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="App强制升级版本" path="appOpen">
                <n-switch
                  v-model:value="formValue.appforce"
                  size="large"
                  @update:value="systemOpenChange2"
                />
              </n-form-item>
            </div>
          </n-gi>
          <n-gi style="margin-top: 6px">
            <div style="margin-left: 40px">
              <n-form-item label="App是否审核" path="appApply">
                <n-switch
                  v-model:value="formValue.appApply"
                  size="large"
                  @update:value="systemOpenChange3"
                />
              </n-form-item>
            </div>
          </n-gi> -->
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

const formBtnLoading = ref(false);
const phoneModelArr = ref([])
const appVersionArr = ref([
  {
    name: 'ios',
    info: {
      app_version: '1.0.0',
      app_privacy: '',
      app_download:'',
      app_version_content:'',
      app_open:false,
      app_audit:false,
      app_forced_upload:false,
      app_pay_model:[],
      minipaychannel:'',
      app_privacy_zh:'',
      app_privacy_zh_CN:'',
      app_privacy_en:'',
      app_privacy_ja:'',
      app_privacy_ko:'',
    }
  }
])
const tabValue = ref('ios')
const group = ref('app');
const show = ref(false);
const formRef: any = ref(null);
const message = useMessage();
const dialog = useDialog();

const formValue = ref({
  androidDownloadUrl: '',
  appOpen: true,
  appforce: false,
  appApply: true,
  appversion: 'v1.0.0',
  appversionContent: '',
  iosDownloadUrl: '',
});

const rules = {
  basicName: {
    required: true,
    message: '请输入网站名称',
    trigger: 'blur',
  },
};

function systemOpenChange(value) {
  dialog.warning({
    title: '提示',
    content: '您确定要app开启或关闭访问吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      if(formValue.value.appOpen){
        formValue.value.appOpen = false;
      }else{
        formValue.value.appOpen = true;
      }
    },
  });
}

function systemOpenChange2(value) {
  dialog.warning({
    title: '提示',
    content: '您确定要app强制升级或取消版本吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      if(formValue.value.appforce){
        formValue.value.appforce = false;
      }else{
        formValue.value.appforce = true;
      }
    },
  });
}

function systemOpenChange3(value) {
  dialog.warning({
    title: '提示',
    content: '您确定要变更app审核状态吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      if(formValue.value.appApply){
        formValue.value.appApply = false;
      }else{
        formValue.value.appApply = true;
      }
    },
  });
}


function versionOpenChange(value) {
  // console.log("value",value);
  dialog.warning({
    title: '提示',
    content: '您确定要app开启或关闭访问吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      // formValue.value.appOpen = true;
    },
  });
}

function versionOpenChange2(value) {
  dialog.warning({
    title: '提示',
    content: '您确定要app强制升级或取消版本吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      // formValue.value.appforce = true;
    },
  });
}

function versionOpenChange3(value) {
  dialog.warning({
    title: '提示',
    content: '您确定要变更app审核状态吗？该操作保存后立马生效，请慎重操作！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
    },
    onNegativeClick: () => {
      // if(formValue.value.appApply){
      //   formValue.value.appApply = false;
      // }else{
      //   formValue.value.appApply = true;
      // }
    },
  });
}

function formSubmit() {
  formBtnLoading.value = true;
  formValue.value.appversion = JSON.stringify(appVersionArr.value)
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

function handleUpdateVersion(values: string[]) {
  phoneModelArr.value = values;
  let values_length = values.length;
  let model_length = appVersionArr.value.length;
  if(model_length < values_length){
    appVersionArr.value.push({
      name: values[values_length - 1],
      info: {
        app_version: '1.0.0',
        app_privacy: '',
      }
    })
  }else if(model_length > values_length){
    let diffIndexes = [];
    appVersionArr.value.forEach((item, index) => {
      if (values.indexOf(item.name) === -1) {
        diffIndexes.push(index);
      }
    });
    // console.log("diffIndexes",diffIndexes)
    diffIndexes.forEach(index => {
      appVersionArr.value.splice(index, 1);
    });
  }

}


onMounted(() => {
  load();
});

function load() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    getConfig({group: group.value})
      .then((res) => {
        formValue.value = res.list;
        let arr = JSON.parse(formValue.value.appversion);
        tabValue.value = arr[0]['name'];
        appVersionArr.value = arr;

        phoneModelArr.value = arr.map(item => String(item.name));
      })
      .finally(() => {
        show.value = false;
      });
  });
}
</script>
