<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="180" :model="formValue" :rules="rules"
              label-placement="left">

        <n-form-item label="酒店订单过期时间" path="smtpPort">
          <n-input-number v-model:value="formValue.hotelStayExp" :show-button="false"
                          placeholder="">
            <template #suffix> 秒</template>
          </n-input-number>
        </n-form-item>

        <n-tabs v-model:value="tabName" size="small" type="card">
          <n-tab-pane name="PayCloud">
            <template #tab> PayCloud</template>
            <n-divider title-placement="left">PayCloud</n-divider>
            <n-form-item label="应用ID">
              <n-input v-model:value="formValue.payCloudAppID" placeholder=""/>
            </n-form-item>
            <n-form-item label="子应用ID">
              <n-input v-model:value="formValue.payCloudSubAppId" placeholder=""/>
            </n-form-item>
            <n-form-item label="微信小程序子应用ID">
              <n-input v-model:value="formValue.payCloudWxMiniAppid" placeholder=""/>
            </n-form-item>
            <n-form-item label="商户号">
              <n-input v-model:value="formValue.payCloudMerchantNo" placeholder=""/>
            </n-form-item>
            <n-form-item label="门店号">
              <n-input v-model:value="formValue.payCloudStoreNo" placeholder=""/>
            </n-form-item>
            <n-form-item label="微信小程序门店号">
              <n-input v-model:value="formValue.payCloudWxMiniStoreNo" placeholder=""/>
            </n-form-item>
            <n-form-item label="私钥">
              <n-input v-model:value="formValue.payCloudPrivateKey" placeholder=""
                       show-password-on="click" type="password"/>
            </n-form-item>
            <n-form-item label="域名">
              <n-input v-model:value="formValue.payCloudEndpoint" placeholder=""/>
            </n-form-item>
            <n-form-item label="支付通知的回调地址">
              <n-input v-model:value="formValue.payCloudNotifyUrl" placeholder=""/>
            </n-form-item>
            <n-form-item label="退款通知的回调地址">
              <n-input v-model:value="formValue.payCloudRefundNotifyUrl" placeholder=""/>
            </n-form-item>
            <n-form-item label="重定向用户页面的URL">
              <n-input v-model:value="formValue.payCloudReturnUrlApp" placeholder=""/>
            </n-form-item>
            <n-form-item label="重定向用户页面的URL">
              <n-input v-model:value="formValue.payCloudReturnUrlWeb" placeholder=""/>
            </n-form-item>
          </n-tab-pane>

          <n-tab-pane name="Paypal">
            <template #tab> Paypal</template>
            <n-divider title-placement="left">Paypal</n-divider>
            <n-form-item label="客户端ID">
              <n-input v-model:value="formValue.payPaypalClientID" placeholder=""/>
            </n-form-item>
            <n-form-item label="密钥">
              <n-input v-model:value="formValue.payPaypalSecret" placeholder=""
                       show-password-on="click" type="password"/>
            </n-form-item>
            <n-form-item label="成功重定向用户页面的URL">
              <n-input v-model:value="formValue.payPaypalReturnUrlApp" placeholder=""/>
            </n-form-item>
            <n-form-item label="取消重定向用户页面的URL">
              <n-input v-model:value="formValue.payPaypalCleanUrlApp" placeholder=""/>
            </n-form-item>
            <n-form-item label="成功重定向用户APP的URL">
              <n-input v-model:value="formValue.payPaypalReturnApp" placeholder=""/>
            </n-form-item>
            <n-form-item label="取消重定向用户APP的URL">
              <n-input v-model:value="formValue.payPaypalCleanApp" placeholder=""/>
            </n-form-item>
            <n-form-item label="IsProd是否是测试环境">
              <n-switch v-model:value="formValue.payPaypalIsProd"/>
            </n-form-item>
          </n-tab-pane>

          <n-tab-pane name="Stripe">
            <template #tab>Stripe</template>
            <n-divider title-placement="left">Stripe</n-divider>
            <n-form-item label="密钥">
              <n-input v-model:value="formValue.payStripeKey" placeholder=""
                       show-password-on="click" type="password"/>
            </n-form-item>
            <n-form-item label="签名密钥">
              <n-input v-model:value="formValue.payStripeSignKey" placeholder=""
                       show-password-on="click" type="password"/>
            </n-form-item>
            <n-form-item label="成功重定向用户页面的URL">
              <n-input v-model:value="formValue.payStripeSuccessURL" placeholder=""/>
            </n-form-item>
            <n-form-item label="取消重定向用户页面的URL">
              <n-input v-model:value="formValue.payStripeCancelURL" placeholder=""/>
            </n-form-item>
            <n-form-item label="成功重定向用户APP的URL">
              <n-input v-model:value="formValue.payStripeSuccessAPP" placeholder=""/>
            </n-form-item>
            <n-form-item label="取消重定向用户APP的URL">
              <n-input v-model:value="formValue.payStripeCancelAPP" placeholder=""/>
            </n-form-item>
          </n-tab-pane>

          <n-tab-pane name="MLILIFE小程序支付">
            <template #tab>MLILIFE小程序支付</template>
            <n-divider title-placement="left">小程序支付</n-divider>
            <n-form-item label="签名密钥">
              <n-input v-model:value="formValue.payMiniSignKey" placeholder=""/>
            </n-form-item>
            <n-form-item label="PID">
              <n-input v-model:value="formValue.payMiniPid" placeholder="" show-password-on="click"
                       type="password"/>
            </n-form-item>
            <n-form-item label="AppId">
              <n-input v-model:value="formValue.payMiniAppid" placeholder="" show-password-on="click"
                       type="password"/>
            </n-form-item>
            <n-form-item label="异步通知URL">
              <n-input v-model:value="formValue.payMiniNotifyUrl" placeholder=""/>
            </n-form-item>
          </n-tab-pane>
        </n-tabs>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
            <!--            <n-button type="default" @click="sendTest">测试支付</n-button>-->
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {useMessage} from 'naive-ui';
import {getConfig, sendTestSms, updateConfig} from '@/api/sys/config';
import {Dicts} from '@/api/dict/dict';
import {Options} from '@/utils/hotgo';

const group = ref('pay');
const show = ref(false);
const showModal = ref(false);
const formBtnLoading = ref(false);
const formParams = ref({mobile: '', event: '', code: '1234'});
const rules = {};
const formTestRef = ref<any>();
const formRef: any = ref(null);
const message = useMessage();

/** 默认选项卡 */
const defaultTabName = 'PayCloud';
/** 选项卡名称 */
const tabName = ref<string>(defaultTabName);

const options = ref<Options>({
  config_sms_template: [],
  config_sms_drive: [],
});


const formValue = ref({
  payCloudAppID: "",
  payCloudSubAppId: "",
  payCloudWxMiniAppid: "",
  payCloudMerchantNo: "",
  payCloudStoreNo: "",
  payCloudWxMiniStoreNo: "",
  payCloudPrivateKey: "",
  payCloudEndpoint: "",
  payCloudNotifyUrl: "",
  payCloudReturnUrlWeb: "",
  payCloudReturnUrlApp: "",
  payStripeKey: "",
  payStripeSignKey: "",
  payStripeSuccessURL: "",
  payStripeCancelURL: "",
  payStripeSuccessAPP: "",
  payStripeCancelAPP: "",
  payPaypalClientID: "",
  payPaypalSecret: "",
  payPaypalReturnUrlApp: "",
  payPaypalCleanUrlApp: "",
  payPaypalReturnApp: "",
  payPaypalCleanApp: "",
  payPaypalIsProd: false,
  hotelStayExp: 600,
  payMiniSignKey:"",
  payMiniPid:"",
  payMiniAppid:"",
  payMiniNotifyUrl:"",
});

function sendTest() {
  showModal.value = true;
  formBtnLoading.value = false;
}

function formSubmit() {
  formRef.value.validate((errors) => {
    if (!errors) {
      updateConfig({group: group.value, list: formValue.value}).then((_res) => {
        message.success('更新成功');
        load();
      });
    } else {
      message.error('验证失败，请填写完整信息');
    }
  });
}

onMounted(() => {
  load();
});

async function load() {
  show.value = true;
  await loadOptions();
  new Promise((_resolve, _reject) => {
    getConfig({group: group.value})
      .then((res) => {
        formValue.value = res.list;
      })
      .finally(() => {
        show.value = false;
      });
  });
}

async function loadOptions() {
  options.value = await Dicts({
    types: ['config_sms_template', 'config_sms_drive'],
  });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formTestRef.value.validate((errors) => {
    if (!errors) {
      sendTestSms(formParams.value).then((_res) => {
        message.success('发送成功');
        showModal.value = false;
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}
</script>
