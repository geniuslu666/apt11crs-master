<template>
  <div class="member-admin-page">
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        :header-style="{
                    padding: '25px 20px 20px',
                  }"
        :content-style="{
                    padding: '0 20px 20px',
                  }"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">基础配置</text>
        </template>
        <n-form :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules" ref="formRef">
          <n-grid cols="1">
            <n-gi>
              <div>
                <n-form-item label-placement="left" path="linkPhone" label="联系电话">
                  <n-input placeholder="请输入联系电话" v-model:value="formValue.linkPhone" style="width: 300px" />
                </n-form-item>
              </div>
            </n-gi>
            <n-gi>
              <div>
                <n-form-item label-placement="left" path="isOpen" label="微信号">
                  <n-input placeholder="请输入微信号" v-model:value="formValue.wechatNo" style="width: 300px" />
                </n-form-item>
              </div>
            </n-gi>
            <n-gi>
              <div>
                <n-form-item label="微信二维码" path="wechatCodeImg">
                  <FileChooser1 v-model:value="formValue.wechatCodeImg" :maxNumber="1"
                                fileType="default"/>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi>
              <div>
                <n-form-item label-placement="left" path="LineId" label="LINE ID">
                  <n-input placeholder="请输入LINE ID" v-model:value="formValue.lineId" style="width: 300px" />
                </n-form-item>
              </div>
            </n-gi>
            <n-gi>
              <div>
                <n-form-item label="LINE图片" path="lineImg">
                  <FileChooser1 v-model:value="formValue.lineImg" :maxNumber="1"
                                fileType="default"/>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi>
              <div>
                <n-form-item label-placement="left" path="isOpen" label="是否发送短信">
                  <n-switch v-model:value="formValue.isSendSms" :unchecked-value="0" :checked-value="1" />
                  <template #feedback>
                    <div style="font-size: 12px">如开启，会员在提交意向后，将会给相应人员发放投资意向短信</div>
                  </template>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi v-if="formValue.isSendSms == 1" style="margin-top: 24px">
              <div>
                <n-form-item label="发送短信电话" path="phonesArr">
                  <n-dynamic-tags v-model:value="phonesArr" />
                  <template #feedback>
                    <div style="font-size: 12px">电话格式：区号-电话， 如：86-18888888888</div>
                  </template>
                </n-form-item>
              </div>
            </n-gi>

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

  const group = ref('member_intention');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const phonesArr = ref([])

  const formValue = ref({
    homepageAdvBanner: '',
    linkPhone: '',
    wechatNo: '',
    wechatCodeImg: '',
    lineId: '',
    lineImg: '',
    isSendSms: 0,
    sendSmsPhone: ''
  });

  const rules = {

  };

  function formSubmit() {
    if(formValue.value.isSendSms != 1){
      phonesArr.value = []
    }
    formValue.value.sendSmsPhone = phonesArr.value.join(",")
    // console.log('formValue.value', formValue.value)
    // return false;
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
          if(res.list.sendSmsPhone!=""){
            phonesArr.value = res.list.sendSmsPhone.split(",").map(item => String(item));
          }
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
