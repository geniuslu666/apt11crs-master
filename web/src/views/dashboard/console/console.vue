<template>
  <div class="p-2">
    <div class="flex-row mt-5 mb-5">
      <div class="flex-item">
        <span class="c333" style="font-size: 45px">
          {{ translang('欢迎') }}，{{ namegetlang(userStore.userPms.nameLanguage) }}!
        </span>
      </div>
      <div style="padding-top: 40px">
        <n-date-picker
          v-model.value="state.date"
          :placeholder="translang('请选择日期')"
          @change="Load()"
          style="width: 180px"
        />
      </div>
    </div>
    <div class="bgfff mb-3 p-5" style="border-radius: 5px">
      <div class="f18 fw c333 mb-2">{{ translang('活动视图') }}</div>
      <n-grid x-gap="12" :cols="4" gutter="12">
        <n-gi>
          <div
            class="grid flex-row"
            style="
              background: linear-gradient(to bottom, #e4e8f1, #f9fafb99);
              position: relative;
              z-index: 0;
            "
          >
            <div style="width: 130px; padding: 15px">
              <n-progress
                :stroke-width="10"
                type="circle"
                :percentage="80"
                :offset-degree="120"
                color="#6883DC"
                indicator-text-color="#333"
                style="width: 100px"
              />
            </div>
            <div class="flex-item lh-2 pt-3 pb-3">
              <div>{{ translang('预抵数') }}</div>
              <div class="f22 fw">{{ state.dashboard.import_bookings }}</div>
              <div class="c12"
                ><span class="c999">{{ translang('已办理入住：') }}</span
                >{{ state.dashboard.checked_bookings }}</div
              >
            </div>
            <img
              style="width: 140px; height: 137px; position: absolute; right: 0; z-index: 99"
              src="@/assets/images/yds.png"
            />
          </div>
        </n-gi>
        <n-gi>
          <div
            class="grid flex-row"
            style="
              background: linear-gradient(to bottom, #deeaff, #f9fafb99);
              position: relative;
              z-index: 0;
            "
          >
            <div style="width: 130px; padding: 15px">
              <n-progress
                :stroke-width="10"
                type="circle"
                :percentage="80"
                :offset-degree="120"
                color="#3BA1FF"
                indicator-text-color="#333"
                style="width: 100px"
              />
            </div>
            <div class="flex-item lh-2 pt-3 pb-3">
              <div>{{ translang('预离数') }}</div>
              <div class="f22 fw">{{ state.dashboard.export_bookings }}</div>
              <div class="c12"
                ><span class="c999">{{ translang('已办理退房：') }}</span
                >{{ state.dashboard.checkout_bookings }}</div
              >
            </div>
            <div class="text-r">
              <img
                style="width: 140px; height: 137px; position: absolute; right: 0; z-index: 99"
                src="@/assets/images/yls.png"
              />
            </div>
          </div>
        </n-gi>
        <n-gi>
          <div
            class="grid flex-row"
            style="
              background: linear-gradient(to bottom, #d8f1ff, #f7fbff);
              position: relative;
              z-index: 0;
            "
          >
            <div style="width: 130px; padding: 15px">
              <n-progress
                :stroke-width="10"
                type="circle"
                :percentage="80"
                :offset-degree="120"
                color="#5DDECF"
                indicator-text-color="#333"
                style="width: 100px"
              />
            </div>
            <div class="flex-item lh-2 pt-3 pb-3">
              <div>{{ translang('在住客房数') }}</div>
              <div class="f22 fw">{{ state.dashboard.checked_bookings }}</div>
              <div class="c12"
                ><span class="c999">{{ translang('客房总数：') }}</span
                >{{ state.dashboard.all_room }}</div
              >
            </div>
            <div class="text-r">
              <img
                style="width: 140px; height: 137px; position: absolute; right: 0; z-index: 99"
                src="@/assets/images/zks.png"
              />
            </div>
          </div>
        </n-gi>
        <n-gi>
          <div
            class="grid flex-row"
            style="background: linear-gradient(to bottom, #f6e8f0, #fbf9fa)"
          >
            <div class="flex-item lh-2 pl-5 pt-3">
              <div>{{ translang('今日预定') }}</div>
              <div class="fw" style="font-size: 35px">{{ state.dashboard.today_stays }}</div>
            </div>
            <div class="text-r">
              <img style="width: 140px; height: 137px" src="@/assets/images/jryd.png" />
            </div>
          </div>
        </n-gi>
      </n-grid>
    </div>
    <div>
      <n-grid x-gap="12" :cols="2">
        <n-gi>
          <div class="bgfff mb-1 p-5" style="border-radius: 5px">
            <div class="f18 fw c333 mb-2">{{ translang('自助入住报告') }}</div>
            <n-grid x-gap="12" :cols="3">
              <n-gi>
                <div
                  class="p-2 c333"
                  style="
                    border-radius: 5px;
                    background-color: #f8f9fd;
                    border-top-right-radius: 50px;
                  "
                >
                  <div class="flex-row mb-3">
                    <div class="pt-2 mr-2">
                      <img style="width: 43px; height: 43px" src="@/assets/images/sz.png" />
                    </div>
                    <div class="flex-item">
                      <div class="fw mb-1" style="font-size: 20px">暂无数据</div>
                      <div class="c333 f12">{{ translang('已提前入住') }}</div>
                    </div>
                  </div>
                  <div class="f12 mb-2 mt-5"
                    ><span class="c999">{{ translang('入住总数：') }}</span
                    >暂无数据</div
                  >
                  <n-progress
                    type="line"
                    :percentage="0"
                    :indicator-placement="'inside'"
                    processing
                    color="#3B7CFF"
                  />
                </div>
              </n-gi>
              <n-gi>
                <div
                  class="p-2 c333"
                  style="
                    border-radius: 5px;
                    background-color: #f8f9fd;
                    border-top-right-radius: 50px;
                  "
                >
                  <div class="flex-row mb-3">
                    <div class="pt-2 mr-2">
                      <img style="width: 43px; height: 43px" src="@/assets/images/rz.png" />
                    </div>
                    <div class="flex-item">
                      <div class="fw mb-1" style="font-size: 20px">暂无数据</div>
                      <div class="c333 f12">{{ translang('已入住') }}</div>
                    </div>
                  </div>
                  <div class="f12 mb-2 mt-5"
                    ><span class="c999">{{ translang('入住总数：') }}</span
                    >暂无数据</div
                  >
                  <n-progress
                    type="line"
                    :percentage="0"
                    :indicator-placement="'inside'"
                    processing
                    color="#5DDECF"
                  />
                </div>
              </n-gi>
              <n-gi>
                <div
                  class="p-2 c333"
                  style="
                    border-radius: 5px;
                    background-color: #f8f9fd;
                    border-top-right-radius: 50px;
                  "
                >
                  <div class="flex-row mb-3">
                    <div class="pt-2 mr-2">
                      <img style="width: 43px; height: 43px" src="@/assets/images/tf.png" />
                    </div>
                    <div class="flex-item">
                      <div class="fw mb-1" style="font-size: 20px">暂无数据</div>
                      <div class="c333 f12">{{ translang('已退房') }}</div>
                    </div>
                  </div>
                  <div class="f12 mb-2 mt-5"
                    ><span class="c999">{{ translang('退房总数：') }}</span
                    >暂无数据</div
                  >
                  <n-progress
                    type="line"
                    :percentage="0"
                    :indicator-placement="'inside'"
                    processing
                    color="#FF5B71"
                  />
                </div>
              </n-gi>
            </n-grid>
          </div>
        </n-gi>
        <n-gi>
          <div class="bgfff mb-1 p-5" style="border-radius: 5px">
            <div class="f18 fw c333 mb-2">{{ translang('清洁任务') }}</div>
            <n-grid x-gap="12" :cols="2">
              <n-gi>
                <div
                  class="text-c p-3 fw c333 mb-2"
                  style="
                    border-radius: 10px;
                    background-color: #f8f9fd;
                    position: relative;
                    font-size: 28px;
                  "
                >
                  <div
                    style="
                      border-radius: 3px;
                      border-top-left-radius: 10px;
                      padding: 4px 12px;
                      background-color: #ffe7e9;
                      color: #d1243b;
                      font-size: 12px;
                      position: absolute;
                      top: 0px;
                      left: 0px;
                      font-weight: 400;
                    "
                    >{{ translang('未分配') }}
                  </div>
                  <span style="font-size: 20px">暂无数据</span>
                </div>
              </n-gi>
              <n-gi>
                <div
                  class="text-c p-3 fw c333 mb-2"
                  style="
                    border-radius: 10px;
                    background-color: #f8f9fd;
                    position: relative;
                    font-size: 28px;
                  "
                >
                  <div
                    style="
                      border-radius: 3px;
                      border-top-left-radius: 10px;
                      padding: 4px 12px;
                      background-color: #ffeedc;
                      color: #d97308;
                      font-size: 12px;
                      position: absolute;
                      top: 0px;
                      left: 0px;
                      font-weight: 400;
                    "
                    >{{ translang('待定') }}
                  </div>
                  <span style="font-size: 20px">暂无数据</span>
                </div>
              </n-gi>
              <n-gi>
                <div
                  class="text-c p-3 fw c333 mb-2"
                  style="
                    border-radius: 10px;
                    background-color: #f8f9fd;
                    position: relative;
                    font-size: 28px;
                  "
                >
                  <div
                    style="
                      border-radius: 3px;
                      border-top-left-radius: 10px;
                      padding: 4px 12px;
                      background-color: #e7f7f6;
                      color: #189f90;
                      font-size: 12px;
                      position: absolute;
                      top: 0px;
                      left: 0px;
                      font-weight: 400;
                    "
                    >{{ translang('准备好清洁') }}
                  </div>
                  <span style="font-size: 20px">暂无数据</span>
                </div>
              </n-gi>
              <n-gi>
                <div
                  class="text-c p-3 fw c333 mb-2"
                  style="
                    border-radius: 10px;
                    background-color: #f8f9fd;
                    position: relative;
                    font-size: 28px;
                  "
                >
                  <div
                    style="
                      border-radius: 3px;
                      border-top-left-radius: 10px;
                      padding: 4px 12px;
                      background-color: #e6eeff;
                      color: #185ce5;
                      font-size: 12px;
                      position: absolute;
                      top: 0px;
                      left: 0px;
                      font-weight: 400;
                    "
                    >{{ translang('确认') }}
                  </div>
                  <span>暂无数据</span>
                </div>
              </n-gi>
            </n-grid>
          </div>
        </n-gi>
      </n-grid>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import { onMounted, reactive, watch } from 'vue';
  import { dashboard } from '@/api/comm';
  import { formatToDateTime } from '@/utils/dateUtil';
  import { getlang, translang } from '@/utils/smjcomm';
  import { useUserStore } from '@/store/modules/user';
  const userStore = useUserStore();
  const state = reactive({
    dashboard: {},
    date: new Date(),
  });
  const Load = async () => {
    const res = await dashboard({
      date: formatToDateTime(state.date, 'yyyy-MM-dd'),
      puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
    });
    state.dashboard = res.Details;
  };
  const namegetlang = (data) => {
    if (data) {
      return getlang(data, userStore.language).content;
    } else {
      return '暂无物业名';
    }
  };
  //页面刷新  根据物业选择
  watch(
    () => userStore.getuserPms,
    () => {
      Load(); //当前页面重新加载的数据
    },
    {
      immediate: true,
      deep: true,
    }
  );
  onMounted(() => {
    Load();
  });
</script>

<style lang="less" scoped>
  .grid {
    width: 100%;
    border-radius: 10px;
    height: 137px;
  }
</style>
