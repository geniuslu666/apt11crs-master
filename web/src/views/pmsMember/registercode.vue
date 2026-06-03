<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="发送会员验证码"
      >
        <n-form label-width="160" label-align="left" :model="formValue" :rules="rules" ref="formRef">
          <n-grid cols="1" :y-gap="10">
            <n-gi>
              <div style="margin-left: 40px">

                <n-form-item label="发送方式" path="goodsState">
                  <n-radio-group v-model:value="formValue.type" name="type">
                    <n-radio-button
                      :value="1"
                      label="短信"
                    />
                    <n-radio-button
                      :value="2"
                      label="邮箱"
                    />
                  </n-radio-group>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi v-if="formValue.type == 1">
              <div style="margin-left: 40px">
                <n-form-item label="手机号码" path="phone" :show-require-mark="true">
                  <n-input-group>
                    <n-input placeholder="区号" v-model:value="formValue.areaNo" style="width: 50px" />
                    <n-input-group-label>-</n-input-group-label>
                    <n-input placeholder="手机号" v-model:value="formValue.phone" style="width: 150px" />
                  </n-input-group>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi v-if="formValue.type == 2">
              <div style="margin-left: 40px">
                <n-form-item label="邮箱" path="email">
                  <n-input v-model:value="formValue.email" placeholder="请输入邮箱" style="width: 230px" />
                </n-form-item>
              </div>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">发送</n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { SendRegisterCode } from '@/api/pmsMember';

  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);

  const formValue = ref({
    type: 1,
    phone: '',
    areaNo: '',
    email: '',
  });

  const rules = {
    basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },
  };

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        SendRegisterCode(formValue.value).then((_res) => {
          formBtnLoading.value = false;
          message.success('发送成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(async () => {
    show.value = true;
    await load();
    show.value = false;
  });

  async function load() {

  }

</script>
