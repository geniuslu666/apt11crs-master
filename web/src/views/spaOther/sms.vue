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
                      <div style="font-size: 12px">如开启，会员在下单后，将会给相应人员发放下单提醒短信</div>
                    </template>
                  </n-form-item>
                </div>
              </n-gi>
              <n-gi>
                <div style="margin-left: 40px">
                  <n-form-item label="模板内容" path="smsTencentTemplate">
<!--                    {{item.content}}-->
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

                    <n-table style="width: 700px">
                      <thead>
                      <tr>
                        <th width="200">服务商</th>
                        <th width="500">联系电话</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr :key="index" v-for="(ispItem, index) in ispList">
                        <td>
                          {{ispItem.name}}
                        </td>
                        <td>
                          <n-dynamic-tags v-model:value="ispItem.sendSmsPhoneArr" :max="3" />
                        </td>
                      </tr>
                      </tbody>
                    </n-table>
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
  import { getConfig, spaSmsConfigUpdate } from '@/api/sys/config';
  import {List as IspList} from "@/api/spaIsp";

  const group = ref('spa_sms_config');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const phonesArr = ref([])
  const ispList = ref([])

  const formValue = ref({
    smsTemplate:[
      {
        smsDrive: "Tencent",
        event: "spaPlaceOrder",
        name: "下单提醒",
        content: "您有一条新的预约订单，请及时处理！",
        isSendSms: 0,
        templateCode: "",
        sendSmsPhoneArr: [],
        sendSmsPhone: "",
        sendSmsPhoneInfo: [
          {
            ispId: {
              ispId: "",
              phone: "",
              phoneArr: []
            },
          }
        ],
      }
    ],
  });

  const rules = {

  };

  function formSubmit() {

    // console.log('formValue.value', formValue.value)
    // console.log('ispList',ispList.value)
    // return false;

    ispList.value.forEach((isp,ispIndex) => {
      if (isp.sendSmsPhoneArr.length > 0){
        ispList.value[ispIndex].sendSmsPhone = isp.sendSmsPhoneArr.join(",");
      }else{
        ispList.value[ispIndex].sendSmsPhone = "";
      }
    })

    formRef.value.validate((errors) => {
      if (!errors) {
        spaSmsConfigUpdate({ group: group.value, list: formValue.value, ispList: ispList.value }).then((_res) => {
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

  async function loadIspList(){
    let IspListOrg = await IspList({
      Pagination: false
    })

    IspListOrg.list.forEach((isp,ispIndex) => {
      if (isp.sendSmsPhone != ""){
        IspListOrg.list[ispIndex].sendSmsPhoneArr = isp.sendSmsPhone.split(",").map(item => String(item));
      }else{
        IspListOrg.list[ispIndex].sendSmsPhoneArr = [];
      }
    })

    ispList.value = IspListOrg.list
  }

  async function load() {
    show.value = true;
    await loadIspList()
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
          formValue.value.smsTemplate = JSON.parse(res.list.smsTemplate)

          // console.log('formValue.value', formValue.value)
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
