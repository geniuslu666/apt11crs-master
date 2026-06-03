<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="分销结算模式" path="appversion">
          <a-radio-group v-model:value="formValue.recommendModel" name="radioGroup">
            <a-radio value="FIRST">最后推荐制</a-radio>
            <a-radio value="LAST">终生推荐制</a-radio>
          </a-radio-group>
        </n-form-item>
        <n-form-item label="全局积分汇率" path="appversion">
          <a-input-number
            v-model:value="formValue.exchangeRate"
            placeholder="全局积分汇率"
            :min="0"
            style="width: 200px"
            :precision="2"
          />
        </n-form-item>
        <n-card style="width: 500px;">
          <n-form-item label="会员邀请注册奖励模式" path="appversion">
            <a-radio-group v-model:value="formValue.memberRegisterAwardModel" name="memberRegisterAwardRadioGroup">
              <a-radio value="points">积分奖励</a-radio>
              <a-radio value="coupon">优惠券奖励</a-radio>
            </a-radio-group>
          </n-form-item>

          <n-form-item label="会员邀请注册奖励" path="memberRegisterAward" v-if="formValue.memberRegisterAwardModel == 'points'">
            <a-input-number
              v-model:value="formValue.memberRegisterAward"
              placeholder="会员邀请注册奖励"
              :min="0"
              style="width: 200px"
              :precision="0"
              addon-after="积分"
            />
          </n-form-item>

          <n-form-item label="会员邀请注册奖励优惠券" path="memberRegisterAward" v-if="formValue.memberRegisterAwardModel == 'coupon'">
            <n-select
              placeholder="请选择优惠券"
              :options="couponList"
              label-field="couponName"
              value-field="id"
              clearable
              filterable
              multiple
              style="width: 350px"
              v-model:value="memberRegisterAwardCouponIdArr"
            />
          </n-form-item>
        </n-card>

        <n-form-item label="渠道邀请注册奖励" path="channelRegisterAward" style="margin-top: 12px">
          <a-input-number
            v-model:value="formValue.channelRegisterAward"
            placeholder="渠道邀请注册奖励"
            :min="0"
            style="width: 200px"
            :precision="0"
          />
        </n-form-item>
        <n-form-item label="会员返佣比例" path="memberBrokerageRate">
          <a-input-number
            v-model:value="formValue.memberBrokerageRate"
            placeholder="会员返佣比例"
            :min="0"
            style="width: 200px"
            :precision="2"
            addon-after="%"
          />
        </n-form-item>
        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {All} from "@/api/pmsCouponType";

  const group = ref('yyconfig');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const couponList = ref([]);
  const memberRegisterAwardCouponIdArr = ref([]);

  const formValue = ref({
    memberExpRate: '',
    recommendModel: '',
    exchangeRate: '',
    memberRegisterAwardModel: '',
    memberRegisterAward: '',
    channelRegisterAward: '',
    memberBrokerageRate: '',
    memberRegisterAwardCoupon: ''
  });

  const rules = {
    basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },
  };

  function formSubmit() {

    console.log('memberRegisterAwardCouponIdArr',memberRegisterAwardCouponIdArr.value)

    formValue.value.memberRegisterAwardCoupon = memberRegisterAwardCouponIdArr.value.join(',')

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
    All({}).then((res) => {
      couponList.value = res.list;
    })
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
          memberRegisterAwardCouponIdArr.value = res.list.memberRegisterAwardCoupon ? res.list.memberRegisterAwardCoupon.split(",").map(item => Number(item)) : [];
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
