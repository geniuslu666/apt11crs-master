<template>
  <div>
    <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <template #header>
        <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">数据查询</text>
      </template>

      <n-form
        ref="formRef"
        inline
        :label-width="100"
        :model="formValue"
        :rules="rules"
        label-placement="left"
        :size="size"
      >
        <n-form-item label="订单号" path="ordersn">
          <n-input v-model:value="formValue.ordersn" placeholder="输入订单号" />
        </n-form-item>
        <n-form-item label="外部订单号" path="outOrderSn">
          <n-input v-model:value="formValue.outOrderSn" placeholder="输入外部订单号" />
        </n-form-item>
        <n-form-item>
          <n-button attr-type="button" @click="formSubmit">
            查询
          </n-button>
        </n-form-item>
      </n-form>
      <n-spin :show="show" style="width: 500px">
        <n-alert :show-icon="false" title="查询结果" type="info" style="width: 500px">
          <div style="margin-left: 10px; line-height: 40px" v-if="!formValue.orderIsEmpty">
            <div>物业名称：{{ formValue.propertyName }}</div>
            <div>下单时间：{{ formValue.orderTime }}</div>
          </div>
          <n-empty description="" v-if="formValue.orderIsEmpty">
            <template #extra>

            </template>
          </n-empty>
        </n-alert>
      </n-spin>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { NTag, useDialog, useMessage } from 'naive-ui';
  import {singleView} from "@/api/pmsRoomReservation";

  const size = ref("medium")
  const show = ref(false);
  const formValue = ref({
    ordersn: '',
    outOrderSn: '',
    propertyName: '',
    orderTime: '',
    orderIsEmpty: true,
  })

  const rules = {
    /*ordersn: {
      required: true,
      message: '请输入酒店订单号',
      trigger: ['input', 'blur']
    }*/
  };

  const message = useMessage();
  const formRef = ref<any>({});

  function formSubmit() {
    formRef.value.validate((errors) => {
      if(formValue.value.ordersn=='' && formValue.value.outOrderSn==''){
        message.error('请输入订单号或外部单号');
        return;
      }
      show.value = true;
      if (!errors) {
        singleView({ ordersn: formValue.value.ordersn, outOrderSn: formValue.value.outOrderSn })
          .then((res) => {
            // console.log(res);
            if (res.id > 0) {
              formValue.value.propertyName = res.propertyDetail.name;
              formValue.value.orderTime = res.createdAt;
              formValue.value.orderIsEmpty = false;
            }else{
              formValue.value.orderIsEmpty = true;
            }
            return;
          })
          .catch((error) => {
            message.error(error.toString());
          })
          .finally(()=>{
            show.value = false;
          });
      } else {
        message.error('查询失败');
      }
    });
  }



</script>

<style lang="less" scoped></style>
