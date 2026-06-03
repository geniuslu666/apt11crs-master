<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">短信配置</text>
        </template>
        <n-form label-width="auto"  label-align="left" :model="formValue" :rules="rules" ref="formRef" label-placement="left" require-mark-placement="right-hanging">
          <n-grid cols="1" :y-gap="1">
            <template v-for="(item, index) of formValue.smsTemplate" :key="index">
              <n-gi >
                <n-divider title-placement="left">{{ formValue.smsTemplate[index].name }}</n-divider>
                <div style="margin-left: 40px">
                  <n-form-item label-placement="left" path="isOpen" label="是否发送短信">
                    <n-switch v-model:value="item.isSendSms" :unchecked-value="0" :checked-value="1" />
                    <template #feedback>
                      <div style="font-size: 12px">{{ item.tips }}</div>
                    </template>
                  </n-form-item>
                </div>
              </n-gi>
              <n-gi>
                <div style="margin-left: 40px">
                  <n-form-item label="模板内容" path="smsTencentTemplate">
                    <n-input v-model:value="item.content" style="width: 350px"
                             placeholder="请输入模板内容" type="textarea"/>
                    <template #feedback>
                      短信平台后台配置
                    </template>
                  </n-form-item>
                </div>
              </n-gi>
              <n-gi>
                <div style="margin-left: 40px">
                  <n-form-item label="模板ID/CODE" path="smsTencentTemplate">
                    <n-input v-model:value="item.templateCode" placeholder="" style="width: 350px" />
                    <template #feedback>

                    </template>
                  </n-form-item>
                </div>
              </n-gi>
              <n-gi v-if="item.isSendSms == 1">
                <div style="margin-left: 40px">
                  <n-form-item label="发送短信电话" path="phonesArr">
                    <n-dynamic-tags v-model:value="item.sendSmsPhoneArr" />
                    <template #feedback>
                      <div style="font-size: 12px">电话格式：区号+电话， 如：+8618888888888</div>
                    </template>
                  </n-form-item>
                </div>
              </n-gi>
            </template>

          </n-grid>
          <div>
            <n-space justify="center">
              <n-button type="primary" @click="formSubmit">保存更新</n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>

    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import FileChooser from '@/components/FileChooser/index.vue';

  const group = ref('car_sms_config');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();

  const formValue = ref({
    smsTemplate:[
      {
        smsDrive: "Tencent",
        event: "carPlaceOrder",
        name: "下单提醒",
        content: "您有一条新的预约订单，请及时处理！",
        isSendSms: 0,
        templateCode: "",
        sendSmsPhoneArr: [],
        sendSmsPhone: "",
      }
    ]
  });

  const rules = {

  };

  function formSubmit() {

    console.log('formValue.value', formValue.value)
    formValue.value.smsTemplate.forEach((item,index) => {
      if(item.isSendSms == 1){
        if(item.sendSmsPhoneArr.length == 0){
          message.error('请填写发送短信电话');
          return false;
        }
      }
      formValue.value.smsTemplate[index].sendSmsPhone = item.sendSmsPhoneArr.join(",")
    })

    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
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

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
          formValue.value.smsTemplate = JSON.parse(res.list.smsTemplate)

          console.log('formValue.value', formValue.value)

          // if( formValue.value.smsTemplate){
          //   formValue.value.smsTemplate.forEach((item,index) => {
          //     formValue.value.smsTemplate[index].sendSmsPhoneArr = item.sendSmsPhone.split(",").map(item => String(item));
          //   })
          // }
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
