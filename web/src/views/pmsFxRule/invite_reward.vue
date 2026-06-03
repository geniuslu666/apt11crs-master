<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form label-width="160" label-align="left" :model="formValue" ref="formRef">
        <n-grid cols="1" :y-gap="20">
          <n-gi>
            <n-divider title-placement="left">
              被邀请人规则
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="isInviteRegOpen" label="是否开启邀请注册奖励">
                <n-switch v-model:value="formValue.isInviteRegOpen" :unchecked-value="0" :checked-value="1" />
                <template #feedback>
                  <div style="font-size: 12px">如开启，新注册的会员在被邀请注册成功后，除了能获得新人注册奖励（如开启），还能获得邀请奖励</div>
                </template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isInviteRegOpen == 1">
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="rewardType" label="奖励类型">
                <n-checkbox-group v-model:value="rewardTypeArr" @update:value="handleUpdateValue">
                  <n-space item-style="display: flex;">
                    <n-checkbox value="balance" label="送积分" />
                    <n-checkbox value="coupon" label="送优惠券" />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isInviteRegOpen == 1 && rewardTypeArr.indexOf('balance') !== -1">
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="rewardBalance" label="奖励积分">
                <a-input-number
                  v-model:value="formValue.rewardBalance"
                  placeholder="建议输入正整数"
                  :min="0"
                  style="width: 200px"
                  :precision="0"
                  addon-after="积分"
                />
                <template #feedback>
                  <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>
                </template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isInviteRegOpen == 1 && rewardTypeArr.indexOf('coupon') !== -1">
            <div style="margin-left: 40px">
              <n-button type="success" @click="chooseCoupon">选择优惠券</n-button>
            </div>
            <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
              <n-table>
                <thead>
                <tr>
                  <th>优惠券名称</th>
                  <th>优惠券类型</th>
                  <th>优惠内容</th>
                  <th>适用场景</th>
                  <th>有效期</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(item, index) of couponTypeArr" :key="index">
                  <td>{{ item.couponName }}</td>
                  <td>
                    <template v-if="item.type == 'reward'">
                      <div :style="{
                        color: '#17A158',
                        width: '38px',
                        height: '22px',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                        border: '1px solid #17A158'
                      }">满减</div>
                    </template>
                    <template v-if="item.type == 'discount'">
                      <div :style="{
                        color: '#F48720',
                        width: '38px',
                        height: '22px',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                        border: '1px solid #F48720'
                      }">折扣</div>
                    </template>
                  </td>
                  <td>
                    <template v-if="item.type == 'reward'">
                      满{{ item.atLeast }}JPY减{{ item.money }}JPY
                    </template>
                    <template v-if="item.type == 'discount'">
                      <p>满{{ item.atLeast }}JPY打{{ item.discount }}折；</p>
                      <p>最多可抵{{item.discountLimit }}JPY;</p>
                    </template>
                  </td>
                  <td>
                    <template v-if="item.scene == 1">
                      <div :style="{
                        display: 'inline-block',
                        color: '#26A763',
                        padding: '0 5px',
                        height: '22px',
                        background: '#E3F4EB',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                      }">民宿</div>
                    </template>
                    <template v-if="item.scene == 2">
                      <div :style="{
                         display: 'inline-block',
                        color: '#3F9EFF',
                        padding: '0 5px',
                        height: '22px',
                        background: '#ECF5FF',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                      }">订餐</div>
                    </template>
                    <template v-if="item.scene == 3">
                      <div :style="{
                         display: 'inline-block',
                        color: '#EFA020',
                        padding: '0 5px',
                        height: '22px',
                        background: '#FFECCE',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                      }">按摩</div>
                    </template>
                    <template v-if="item.scene == 4">
                      <div :style="{
                         display: 'inline-block',
                        color: '#F56C6C',
                        padding: '0 5px',
                        height: '22px',
                        background: '#FEF0F0',
                        lineHeight: '22px',
                        textAlign: 'center',
                        fontSize: '14px',
                        borderRadius: '2px',
                        fontWeight: '400',
                      }">接送机/包车</div>
                    </template>
                  </td>
                  <td>
                    <template v-if="item.validityType == 1">
                      {{ item.endUseTime }}
                    </template>
                    <template v-else-if="item.validityType == 2">
                      领取后{{ item.fixedTerm }}天有效
                    </template>
                    <template v-else>
                      长期有效
                    </template>
                  </td>
                  <td>
                    <n-button type="error" @click="handleDeleteCoupon(index)">删除</n-button>
                  </td>
                </tr>
                </tbody>
              </n-table>
            </div>
          </n-gi>

          <n-gi>
            <n-divider title-placement="left">
              邀请人奖励
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="isInviteRewardOpen" label="是否开启邀请人奖励">
                <n-switch v-model:value="formValue.isInviteRewardOpen" :unchecked-value="0" :checked-value="1" />
                <template #feedback>
                  <div style="font-size: 12px">如开启，邀请人在成功邀请用户注册后，可获得以下配置的奖励</div>
                </template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isInviteRewardOpen == 1">
            <div style="margin-left: 40px">
              <n-form-item label-placement="left" path="isInviteRewardOpen" label="奖励发放阶段:">
                新用户首次订单完成后
                <template #feedback>
                  <div style="font-size: 12px">用户首次订单完成后：当用户受邀注册后，必须完成首次下单且当订单完成后，奖励才发放</div>
                </template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isInviteRewardOpen == 1">
            <div style="margin-left: 40px;display: flex">
              <div style="width: 160px;line-height: 34px">奖励配置</div>
              <div>
                <div style="display: block;margin-bottom: 30px">
                  <n-form-item label-placement="left" path="memberRewardBalance" label="普通会员" label-width="80">
                    <a-input-number
                      v-model:value="formValue.memberRewardBalance"
                      placeholder="建议输入正整数"
                      :min="0"
                      style="width: 200px"
                      :precision="0"
                      addon-after="积分"
                    />
                    <template #feedback>
                      <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>
                    </template>
                  </n-form-item>
                </div>
                <div style="display: block;margin-bottom: 30px">
                  <n-form-item label-placement="left" path="staffRewardBalance" label="员工" label-width="80">
                    <a-input-number
                      v-model:value="formValue.staffRewardBalance"
                      placeholder="建议输入正整数"
                      :min="0"
                      style="width: 200px"
                      :precision="0"
                      addon-after="积分"
                    />
                    <template #feedback>
                      <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>
                    </template>
                  </n-form-item>
                </div>
                <div style="display: block;margin-bottom: 30px">
                  <n-form-item label-placement="left" path="channelRewardBalance" label="渠道" label-width="80">
                    <a-input-number
                      v-model:value="formValue.channelRewardBalance"
                      placeholder="建议输入正整数"
                      :min="0"
                      style="width: 200px"
                      :precision="0"
                      addon-after="积分"
                    />
                    <template #feedback>
                      <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>
                    </template>
                  </n-form-item>
                </div>
              </div>
            </div>
          </n-gi>
<!--          <n-gi v-if="formValue.isInviteRewardOpen == 1">-->
<!--            <div style="margin-left: 40px">-->
<!--              <n-form-item label-placement="left" path="memberRewardBalance" label="普通会员">-->
<!--                <a-input-number-->
<!--                  v-model:value="formValue.memberRewardBalance"-->
<!--                  placeholder="建议输入正整数"-->
<!--                  :min="0"-->
<!--                  style="width: 200px"-->
<!--                  :precision="0"-->
<!--                  addon-after="积分"-->
<!--                />-->
<!--                <template #feedback>-->
<!--                  <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>-->
<!--                </template>-->
<!--              </n-form-item>-->
<!--            </div>-->
<!--          </n-gi>-->
<!--          <n-gi v-if="formValue.isInviteRewardOpen == 1">-->
<!--            <div style="margin-left: 40px">-->
<!--              <n-form-item label-placement="left" path="staffRewardBalance" label="员工">-->
<!--                <a-input-number-->
<!--                  v-model:value="formValue.staffRewardBalance"-->
<!--                  placeholder="建议输入正整数"-->
<!--                  :min="0"-->
<!--                  style="width: 200px"-->
<!--                  :precision="0"-->
<!--                  addon-after="积分"-->
<!--                />-->
<!--                <template #feedback>-->
<!--                  <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>-->
<!--                </template>-->
<!--              </n-form-item>-->
<!--            </div>-->
<!--          </n-gi>-->
<!--          <n-gi v-if="formValue.isInviteRewardOpen == 1">-->
<!--            <div style="margin-left: 40px">-->
<!--              <n-form-item label-placement="left" path="channelRewardBalance" label="渠道">-->
<!--                <a-input-number-->
<!--                  v-model:value="formValue.channelRewardBalance"-->
<!--                  placeholder="建议输入正整数"-->
<!--                  :min="0"-->
<!--                  style="width: 200px"-->
<!--                  :precision="0"-->
<!--                  addon-after="积分"-->
<!--                />-->
<!--                <template #feedback>-->
<!--                  <div style="font-size: 12px">积分必须为整数，0表示不赠送</div>-->
<!--                </template>-->
<!--              </n-form-item>-->
<!--            </div>-->
<!--          </n-gi>-->
        </n-grid>
        <div style="text-align: center">
          <n-space justify="center">
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>

    <ChooseCoupon ref="chooseCouponRef" @reloadCouponList="chooseCouponInfo"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import ChooseCoupon from "@/components/ChooseCouponType/chooseCouponType.vue";
  import {All} from "@/api/pmsCouponType";

  const group = ref('invitenewreward');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const chooseCouponRef = ref();
  const formBtnLoading = ref(false);

  const formValue = ref({
    isInviteRegOpen: 0,
    rewardType: '',
    rewardBalance: 0,
    couponTypeIds: '',
    isInviteRewardOpen: 0,
    memberRewardBalance: 0,
    staffRewardBalance: 0,
    channelRewardBalance: 0,
  });
  const rewardTypeArr = ref([]);
  const couponTypeIdsArr = ref([]);
  const couponTypeArr = ref([])

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

  function  handleUpdateValue(value) {
    formValue.value.rewardType = value.join(',')
  }

  function chooseCoupon(){
    // chooseCouponRef.value.openModal(couponTypeIdsArr.value);
    chooseCouponRef.value.openModal([]);
  }

  function handleDeleteCoupon(index){
    couponTypeArr.value.splice(index, 1);
    couponTypeIdsArr.value.splice(index, 1);
    formValue.value.couponTypeIds = couponTypeIdsArr.value.join(',')
  }

  function chooseCouponInfo(couponInfo){
    couponTypeArr.value.push(couponInfo)
    couponTypeIdsArr.value.push(couponInfo.id)
    formValue.value.couponTypeIds = couponTypeIdsArr.value.join(',')
  }

  onMounted(() => {
    load();
  });

  function getCouponList(){
    All({
      couponIds: formValue.value.couponTypeIds
    }).then((res) => {
      couponTypeArr.value = res.list;
    })
      .finally(() => {
        show.value = false;
      });
  }

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
          rewardTypeArr.value = res.list.rewardType.split(',')
          if(res.list.couponTypeIds){
            couponTypeIdsArr.value = res.list.couponTypeIds.split(',').map((item) => {
              return Number(item)
            })
            console.log(couponTypeIdsArr.value)
            getCouponList()
          }else{
            couponTypeIdsArr.value = []
            show.value = false;
          }
        })
        .catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
