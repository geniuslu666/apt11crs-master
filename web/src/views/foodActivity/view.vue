<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="订餐活动表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" column="1">
            <n-descriptions-item>
              <template #label>
                活动日期
              </template>
              {{ formValue.date }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题-简体中文
              </template>
               {{ formValue.zh_name ?? '--' }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题-English
              </template>
              {{ formValue.en_name ?? '--' }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题-日本语
              </template>
              {{ formValue.ja_name ?? '--' }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题-한국어
              </template>
              {{ formValue.ko_name ?? '--' }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题-繁体中文
              </template>
               {{ formValue.zh_CN_name ?? '--' }}
            </n-descriptions-item>
<!--            <n-descriptions-item>-->
<!--              <template #label>-->
<!--                副标题-->
<!--              </template>-->
<!--              {{ formValue.subName }}-->
<!--            </n-descriptions-item>-->
            <n-descriptions-item>
              <template #label>
                图片
              </template>
              <img
                v-if="formValue.pic"
                :src="formValue.pic"
                loading="lazy"
                style="width: 100px; border-radius: 100%"
                :onerror="picErr"
              />
              <img
                v-else
                src="~@/assets/images/onerror.png"
                style="width: 100px; border-radius: 100%"
              />
            </n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag :type="getOptionTag(options.sys_normal_disable, formValue?.status)" size="small" class="min-left-space">
                {{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="餐厅">
            <div style="margin-left: 40px;margin-top: 10px;margin-bottom: 20px">
              <n-table>
                <thead>
                <tr>
                  <th>餐厅名称</th>
                  <th>营业类型</th>
                  <th>地址信息</th>
                  <th>餐厅状态</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(item, index) of couponTypeArr" :key="index">
                  <td>{{ item.name }}</td>
                  <td>
                    {{ item.cooperateTypeDetail.typeName }}
                  </td>
                  <td>{{ item.detailAddress }}</td>
                  <td>
                    <n-tag :type="getOptionTag(RestaurantOptions.open_status, item?.openStatus)" size="small" class="min-left-space">
                      {{ getOptionLabel(RestaurantOptions.open_status, item?.openStatus) }}
                    </n-tag>
                  </td>
                </tr>
                </tbody>
              </n-table>
            </div>
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/foodActivity';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import defaultImg from "@/assets/images/onerror.png";
import {jsontoobj} from "@/utils/smjcomm";
import {
  loadOptions as RestaurantLoadOptions,
  options as RestaurantOptions
} from "@/views/foodRestaurant/model";
import {All} from "@/api/foodRestaurant";

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});
const couponTypeIdsArr = ref([]);
const couponTypeArr = ref([])

function picErr(e){
  e.target.src = defaultImg
}

function getRestaurantList(){
  All({
    restaurantIds: formValue.value.restaurantIds
  }).then((res) => {
    couponTypeArr.value = res.list;
  })
    .finally(() => {
      loading.value = false;
    });
}

//下载
function download(url: string) {
  window.open(url);
}

function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  RestaurantLoadOptions()
  View({ id: state.id })
    .then((res) => {
      res.nameLanguage = jsontoobj(res.nameLanguage);

      formValue.value = res;

      formValue.value.zh_name = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
      formValue.value.en_name = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
      formValue.value.ko_name = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
      formValue.value.ja_name = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
      if(res.nameLanguage.zh_CN){
        formValue.value.zh_CN_name = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
      }else if(res.nameLanguage.zh_cn){
        formValue.value.zh_CN_name = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
      }else{
        formValue.value.zh_CN_name = ''
      }


      res.restaurantList.forEach((item) => {
        couponTypeIdsArr.value.push(item.restaurantId)
      })
      formValue.value.restaurantIds = couponTypeIdsArr.value.join(',')
      getRestaurantList()
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped></style>


