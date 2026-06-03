<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="200" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              个人中心按钮开关
            </n-divider>
          </n-gi>
          <n-gi v-for="item in buttonList" :key="item.key">
            <div style="margin-left: 40px">
              <n-form-item :label="item.label">
                <n-switch
                  v-model:value="formValue[item.key]"
                  :checked-value="1"
                  :unchecked-value="0"
                />
                <span style="margin-left: 12px; color: #999; font-size: 13px;">{{ item.desc }}</span>
              </n-form-item>
            </div>
          </n-gi>
        </n-grid>
        <div style="text-align: center; margin-top: 20px;">
          <n-space justify="center">
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { getConfig, updateConfig } from '@/api/sys/config';

const formBtnLoading = ref(false);
const group = ref('profileButtons');
const show = ref(false);
const formRef: any = ref(null);
const message = useMessage();

const buttonList = [
  { key: 'qrcode', label: '会员码', desc: '' },
  { key: 'coupon', label: '优惠券', desc: '' },
  { key: 'thCoupon', label: '礼品券', desc: '' },
  { key: 'dataBoard', label: '数据看板', desc: '' },
  { key: 'workbench', label: '工作台', desc: ''},
  { key: 'employee', label: '员工专区', desc: '' },
  { key: 'fx', label: '渠道分销', desc: '' },
  { key: 'help', label: '帮助中心', desc: '' },
  { key: 'contact', label: '联系我们', desc: '' },
  { key: 'agreement', label: '用户协议', desc: '' },
  { key: 'privacy', label: '隐私政策', desc: '' },
];

const formValue = ref<Record<string, number>>(
  Object.fromEntries(buttonList.map((item) => [item.key, 1]))
);

const rules = {};

function formSubmit() {
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      const submitData = {
        profileButtonData: JSON.stringify(formValue.value),
      };
      updateConfig({ group: group.value, list: submitData })
        .then(() => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        })
        .catch(() => {
          formBtnLoading.value = false;
        });
    } else {
      formBtnLoading.value = false;
      message.error('验证失败');
    }
  });
}

onMounted(() => {
  load();
});

function load() {
  show.value = true;
  getConfig({ group: group.value })
    .then((res) => {
      if (res.list && res.list.profileButtonData) {
        try {
          const data = JSON.parse(res.list.profileButtonData);
          for (const item of buttonList) {
            if (data[item.key] !== undefined) {
              formValue.value[item.key] = data[item.key];
            }
          }
        } catch (e) {
          console.error('解析个人中心按钮配置失败', e);
        }
      }
    })
    .finally(() => {
      show.value = false;
    });
}
</script>
