<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form label-width="auto" :model="formValue" ref="formRef">
        <n-grid cols="1" :y-gap="20">
          <n-gi>
            <n-divider title-placement="left">
              分销全局规则
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item path="recommendModel">
                <template #label>
                  <span class="n-form-item-label__text">分销结算模式</span><span style="margin-left: 5px;color: red">*请谨慎修改</span>
                </template>
                <a-radio-group v-model:value="formValue.recommendModel" name="radioGroup">
                  <a-radio value="FIRST">最后推荐制</a-radio>
                  <a-radio value="LAST">终生推荐制</a-radio>
                </a-radio-group>
                <template #feedback>
                  <div style="font-size: 12px">最后推荐制：被推荐人的推荐人信息可能会变，以他最后一个推荐人为准</div>
                  <div style="font-size: 12px">终生推荐制：用户的推荐人固定为首次成功推荐他的人，且不会变化</div>
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
    recommendModel: '',
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
