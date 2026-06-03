<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form label-width="auto" :model="formValue" ref="formRef">
        <n-grid cols="1" :y-gap="20">
          <n-gi>
            <n-divider title-placement="left">
              推荐下单奖励配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="memberBrokerageRate">
                <template #label>
                  <span class="n-form-item-label__text">普通会员推荐用户下单成功奖励</span>
                </template>
                <a-input-number
                  v-model:value="formValue.memberBrokerageRate"
                  placeholder="建议输入正整数"
                  :min="0"
                  style="width: 200px"
                  :precision="0"
                  addon-after="%"
                />
                <template #feedback>
                  <div style="font-size: 12px">普通会员推荐用户在APP完成下单，在订单完成后普通会员获得的奖励比例</div>
                  <div style="font-size: 12px;color: red">*渠道和员工的奖励在渠道/员工管理内配置</div>
                </template>
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
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('yyconfig');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);

  const formValue = ref({
    memberBrokerageRate: ''
  });

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
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

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
