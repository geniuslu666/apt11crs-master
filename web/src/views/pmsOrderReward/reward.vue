<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form label-width="160" label-align="left" :model="formValue" ref="formRef">
          <n-grid cols="1" :y-gap="10">
            <n-gi>
              <div style="margin-left: 40px">
                <n-form-item label-placement="left" path="isOpen" label="是否开启下单奖励">
                  <n-switch v-model:value="formValue.isOpenReward" :unchecked-value="2" :checked-value="1" />
                  <template #feedback>
                    <div style="font-size: 12px">如开启，会员在下单支付成功后都能获得如下配置的奖励</div>
                  </template>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi v-if="formValue.isOpenReward == 1">
              <div style="margin-left: 40px">
                <n-form-item label-placement="left" path="rewardType" label="奖励类型">
                  <n-checkbox-group v-model:value="rewardTypeArr" @update:value="handleUpdateValue">
                    <n-space item-style="display: flex;">
                      <n-checkbox value="coupon" label="送优惠券" />
                      <n-checkbox value="thcoupon" label="送礼品券" />
                    </n-space>
                  </n-checkbox-group>
                </n-form-item>
              </div>
            </n-gi>
            <n-gi v-if="formValue.isOpenReward == 1 && rewardTypeArr.indexOf('coupon') !== -1">
              <div style="margin-left: 40px">
                <n-button type="primary" @click="chooseCoupon">选择优惠券</n-button>
              </div>
              <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
                <n-table>
                  <thead>
                    <tr>
                      <th>优惠券名称</th>
                      <th>优惠券类型</th>
                      <th>优惠券金额/折扣</th>
                      <th>满多少元使用</th>
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
                          <span style="color: green">满减</span>
                        </template>
                        <template v-if="item.type == 'discount'">
                          <span style="color: blue">折扣</span>
                        </template>
                      </td>
                      <td>
                        <template v-if="item.type == 'reward'">
                          {{ item.money }}JPY
                        </template>
                        <template v-if="item.type == 'discount'">
                          {{ item.discount }}折
                        </template>
                      </td>
                      <td>{{ item.atLeast }}</td>
                      <td>
                        <template v-if="item.scene == 1">
                          <span style="color: green">民宿</span>
                        </template>
                        <template v-if="item.scene == 2">
                          <span style="color: blue">餐饮</span>
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

            <n-gi v-if="formValue.isOpenReward == 1 && rewardTypeArr.indexOf('thcoupon') !== -1">
              <div style="margin-left: 40px">
                <n-button type="primary" @click="chooseThCoupon">选择礼品券</n-button>
              </div>
              <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
                <n-table>
                  <thead>
                  <tr>
                    <th>礼品券名称</th>
                    <th>适用商户</th>
                    <th>发行状态</th>
                    <th>使用状态</th>
                    <th>有效期</th>
                    <th>操作</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="(item, index) of thCouponArr" :key="index">
                    <td>{{ item.couponName }}</td>
                    <td>
                      <n-ellipsis expand-trigger="click" :line-clamp="1" :tooltip="false" style="width: 200px">
                        <div v-for="(item1, index1) of item.mchList" :key="index1">{{ item1.mchInfo?.name }}</div>
                      </n-ellipsis>
                    </td>
                    <td>
                      <n-tag :type="getOptionTag(options.th_coupon_status, item.status)">
                        {{ getOptionLabel(options.th_coupon_status, item.status) }}
                      </n-tag>
                    </td>
                    <td>
                      <n-tag :type="getOptionTag(options.th_coupon_use_status, item.useStatus)">
                        {{ getOptionLabel(options.th_coupon_use_status, item.useStatus) }}
                      </n-tag>
                    </td>
                    <td>
                      激活后{{ item.fixedTerm }}天内有效
                    </td>
                    <td>
                      <n-button type="error" @click="handleDeleteThCoupon(index)">删除</n-button>
                    </td>
                  </tr>
                  </tbody>
                </n-table>
              </div>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
    <ChooseCoupon ref="chooseCouponRef" @reloadCouponList="chooseCouponInfo"/>
    <ThChooseCoupon ref="chooseThCouponRef" @reloadThCouponList="chooseThCouponInfo"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { Edit, View } from '@/api/pmsMemberScene';
  import ChooseCoupon from "@/components/ChooseCouponType/chooseCouponType.vue";
  import ThChooseCoupon from "@/components/ThChooseCouponType/chooseThCouponType.vue";
  import {All} from "@/api/pmsCouponType";
  import {All as ThCouponAll} from "@/api/thCoupon";
  import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
  import {options, loadOptions} from "@/views/thCoupon/model";
  import {newState, State} from "@/views/pmsMemberScene/model";

  interface Props {
    sceneId?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    sceneId: '0',
  });
  const message = useMessage();
  const loading = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const chooseCouponRef = ref();
  const chooseThCouponRef = ref();
  const rewardTypeArr = ref(null);
  const couponTypeIdsArr = ref([]);
  const couponTypeArr = ref([])

  const thCouponIdsArr = ref([]);
  const thCouponArr = ref([])

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value ).then((_res) => {
          formBtnLoading.value = false;
          message.success('操作成功');
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

  function handleUpdateValue(value) {
    value = value.filter(item => item !== null && item !== undefined && item !== '')
    formValue.value.rewardType = value.join(',')
  }

  function chooseCoupon(){
    // chooseCouponRef.value.openModal(couponTypeIdsArr.value);
    chooseCouponRef.value.openModal([]);
  }

  function handleDeleteCoupon(index){
    couponTypeArr.value.splice(index, 1);
    couponTypeIdsArr.value.splice(index, 1);
    formValue.value.rewardCouponTypeIds = couponTypeIdsArr.value.join(',')
  }

  function chooseCouponInfo(couponInfo){
    couponTypeArr.value.push(couponInfo)
    couponTypeIdsArr.value.push(couponInfo.id)
    formValue.value.rewardCouponTypeIds = couponTypeIdsArr.value.join(',')
  }

  function chooseThCoupon(){
    // chooseCouponRef.value.openModal(couponTypeIdsArr.value);
    chooseThCouponRef.value.openModal([]);
  }

  function handleDeleteThCoupon(index){
    thCouponArr.value.splice(index, 1);
    thCouponIdsArr.value.splice(index, 1);
    formValue.value.rewardThCouponIds = thCouponIdsArr.value.join(',')
  }

  function chooseThCouponInfo(couponInfo){
    thCouponArr.value.push(couponInfo)
    thCouponIdsArr.value.push(couponInfo.id)
    formValue.value.rewardThCouponIds = thCouponIdsArr.value.join(',')
  }

  onMounted(async () => {
    loading.value = true;
    await loadOptions()
    await load();
    await getCouponList();
    await getThCouponList();
    loading.value = false;
  });

  async function getCouponList(){
    const res = await All({
      couponIds: formValue.value.rewardCouponTypeIds ? formValue.value.rewardCouponTypeIds : '0'
    });
    couponTypeArr.value = res.list;
  }

  async function getThCouponList(){
    const res = await ThCouponAll({
      couponIds: formValue.value.rewardThCouponIds ? formValue.value.rewardThCouponIds : '0'
    });
    thCouponArr.value = res.list;
  }

  async function load() {
    const res =  await View({ id: props.sceneId })
    formValue.value = res;
    rewardTypeArr.value = res.rewardType.split(',')
    if(res.rewardCouponTypeIds){
      couponTypeIdsArr.value = res.rewardCouponTypeIds.split(',').map((item) => {
        return Number(item)
      })
    }else{
      couponTypeIdsArr.value = []
    }

    if(res.rewardThCouponIds){
      thCouponIdsArr.value = res.rewardThCouponIds.split(',').map((item) => {
        return Number(item)
      })
    }else{
      thCouponIdsArr.value = []
    }
  }
</script>
