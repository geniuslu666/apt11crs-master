<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <div class="border-bottom">
        <n-card :bordered="false">
          <template #header>
            <span style="font-size: 24px">{{ translang('物业详情') }}</span>
          </template>
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;display: flex;">
              <div class="img-cover">
                <img v-if="data.cover" :src="data.cover"/>
                <img v-else src="@/assets/images/nowy.png"/>
                <n-upload
                  :action="`${uploadUrl}${urlPrefix}/upload/file`"
                  :headers="uploadHeaders"
                  @finish="finishUploadCover"
                >
                  <a href="javascript:void(0);">
                    <n-icon size="13" color="#FFFFFF">
                      <EditOutlined/>
                    </n-icon>
                    <span>点击编辑</span>
                  </a>
                </n-upload>
                <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;">建议尺寸：192px*282px</div>
              </div>
              <div>
                <div style="flex: 1;margin-left: 30px">
                  <div
                    style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                    {{ translang('物业模式') }}
                  </div>

                  <n-descriptions label-placement="top" content-class="descriptions-cont"
                                  :column="3">
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">开启预订模式</span>
                      </template>
                      <n-switch v-model:value="data.bookingClose" :loading="bookingCloseloadingRef" @update:value="handleChangebookingClose"/>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">开启短租模式</span>
                      </template>
                      <n-switch v-model:value="data.leaseClose" :loading="leaseCloseloadingRef" @update:value="handleChangeleaseClose"/>
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">开启多房间预定</span>
                      </template>
                      <n-switch v-model:value="data.batchReservation" :loading="batchReservationloadingRef" @update:value="handleChangebatchReservation"/>
                    </n-descriptions-item>
                  </n-descriptions>
                </div>
                <div style="flex: 1;margin-left: 30px">
                  <div
                    style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                    {{ translang('物业名称') }}
                  </div>
                  <n-descriptions label-placement="top" content-class="descriptions-cont"
                                  :column="2">
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">简体中文</span>
                      </template>
                      {{
                        data.nameLanguage.zh && data.nameLanguage.zh.content ? data.nameLanguage.zh.content : '--'
                      }}
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">日本语</span>
                      </template>
                      {{
                        data.nameLanguage.ja && data.nameLanguage.ja.content ? data.nameLanguage.ja.content : '--'
                      }}
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">English</span>
                      </template>
                      {{
                        data.nameLanguage.en && data.nameLanguage.en.content ? data.nameLanguage.en.content : '--'
                      }}
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">한국어</span>
                      </template>
                      {{
                        data.nameLanguage.ko && data.nameLanguage.ko.content ? data.nameLanguage.ko.content : '--'
                      }}
                    </n-descriptions-item>
                    <n-descriptions-item>
                      <template #label>
                        <span class="descriptions-label">繁体中文</span>
                      </template>
                      {{
                        data.nameLanguage.zh_CN && data.nameLanguage.zh_CN.content ? data.nameLanguage.zh_CN.content : '--'
                      }}
                    </n-descriptions-item>
                  </n-descriptions>
                </div>
              </div>

            </div>
            <n-button type="primary"
                      @click="editItem('Name',translang('编辑物业名称'),data.nameLanguage)"
                      style="margin-top: 92px">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('物业地址') }}
              </div>
              <n-descriptions label-placement="top" content-class="descriptions-cont" :column="2" title="所属地区" style="margin-bottom: 12px">
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">简体中文</span>
                  </template>
                  {{
                    data.regionNameLanguage.zh && data.regionNameLanguage.zh.content  ? data.regionNameLanguage.zh.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">日本语</span>
                  </template>
                  {{
                    data.regionNameLanguage.ja && data.regionNameLanguage.ja.content ? data.regionNameLanguage.ja.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">English</span>
                  </template>
                  {{
                    data.regionNameLanguage.en && data.regionNameLanguage.en.content ? data.regionNameLanguage.en.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">한국어</span>
                  </template>
                  {{
                    data.regionNameLanguage.ko && data.regionNameLanguage.ko.content ? data.regionNameLanguage.ko.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">繁体中文</span>
                  </template>
                  {{
                    data.regionNameLanguage.zh_CN && data.regionNameLanguage.zh_CN.content ? data.regionNameLanguage.zh_CN.content : '--'
                  }}
                </n-descriptions-item>
              </n-descriptions>
              <n-descriptions label-placement="top" content-class="descriptions-cont" :column="2" title="详细地址">
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">简体中文</span>
                  </template>
                  {{
                    data.addressLanguage.zh && data.addressLanguage.zh.content ? data.addressLanguage.zh.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">日本语</span>
                  </template>
                  {{
                    data.addressLanguage.ja && data.addressLanguage.ja.content ? data.addressLanguage.ja.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">English</span>
                  </template>
                  {{
                    data.addressLanguage.en && data.addressLanguage.en.content ? data.addressLanguage.en.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">한국어</span>
                  </template>
                  {{
                    data.addressLanguage.ko && data.addressLanguage.ko.content ? data.addressLanguage.ko.content : '--'
                  }}
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">繁体中文</span>
                  </template>
                  {{
                    data.addressLanguage.zh_CN && data.addressLanguage.zh_CN.content ? data.addressLanguage.zh_CN.content : '--'
                  }}
                </n-descriptions-item>
              </n-descriptions>
            </div>
            <n-button type="primary"
                      @click="editItem('Address',translang('编辑物业地址'),data.addressLanguage, data.regionId)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div>
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('经纬度') }}
              </div>
              <div class="map-location-info">
                <span>{{
                    translang('经度')
                  }}：<span>{{
                      mapview == 'baidu' ? (data.lat ? (data.lat == 'undefined' ? '无' : data.lat) : '无') : (data.ggLat ? (data.ggLat == 'undefined' ? '无' : data.ggLat) : '无')
                    }}</span></span>
                <span>{{
                    translang('纬度')
                  }}：<span>{{
                      mapview == 'baidu' ? (data.lng ? (data.lng == 'undefined' ? '无' : data.lng) : '无') : (data.ggLng ? (data.ggLng == 'undefined' ? '无' : data.ggLng) : '无')
                    }}</span></span>
              </div>
            </div>
            <n-button type="primary" v-if="!editMap" @click="handleEditGg">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
            <n-button type="primary" v-else @click="editMap = false">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('退出编辑') }}
            </n-button>
          </div>
          <div style="position: relative">
            <!--            <div class="flex-row mapdiv">-->
            <!--              &lt;!&ndash; 百度 切换 &ndash;&gt;-->
            <!--              <div class="tabview mr-4" style="border-right: 1px solid #e2e2e2; padding-right: 15px">-->
            <!--                <img src="@/assets/images/baidu.png" />-->
            <!--                <span @click="mapview = 'baidu'" :class="mapview == 'baidu' ? 'showmap' : ''">{{-->
            <!--                    translang('百度地图')-->
            <!--                  }}</span>-->
            <!--              </div>-->
            <!--              &lt;!&ndash; 谷歌切换 &ndash;&gt;-->
            <!--              <div class="tabview">-->
            <!--                <img src="@/assets/images/google.png" />-->
            <!--                <span @click="mapview = 'google'" :class="mapview == 'google' ? 'showmap' : ''">{{-->
            <!--                    translang('谷歌地图')-->
            <!--                  }}</span>-->
            <!--              </div>-->
            <!--            </div>-->
            <div v-if="!editMap">
              <div v-if="mapview == 'baidu'">
                <!-- 地址详细信息 -->
                <div class="mapbtm f12" v-if="data.addressDetail">
                  {{ data.addressDetail }}
                </div>
                <div class="mapbtm f12" v-else> {{ translang('暂无定位信息') }}</div>

                <img
                  style="width: 100%; height: 260px; object-fit: cover"
                  class="mt-3 mb-5"
                  src="@/assets/images/nomap.png"
                  v-if="!data.lat || data.lat == 'undefined'"
                />

                <BMapView
                  class="mt-3"
                  v-else
                  :key="mapkey"
                  :latitude="data.lat"
                  :longitude="data.lng"
                  :heightmap="260"
                />
              </div>
              <div v-else>
                <!-- 地址详细信息 -->
                <div class="mapbtm f12" v-if="data.ggAddressDetail">
                  {{ data.ggAddressDetail }}
                </div>
                <div class="mapbtm f12" v-else> {{ translang('暂无定位信息') }}</div>

                <img
                  style="width: 100%; height: 260px; object-fit: cover"
                  class="mt-3 mb-5"
                  src="@/assets/images/nomap.png"
                  v-if="!data.ggLat || data.ggLat == 'undefined'"
                />
                <GMapView
                  class="mt-3"
                  :lat="data.ggLat"
                  :long="data.ggLng"
                  :heightmap="260"
                  v-else
                />
              </div>
            </div>

            <div v-else class="mt-3">
              <baiduMap
                v-if="mapview == 'baidu'"
                :key="mapkey"
                :latitude="data.lat"
                :longitude="data.lng"
                :heightmap="260"
                @addressData="addressDataLoad($event)"
              />
              <googleMap
                v-else
                :key="mapkey"
                :lat="data.ggLat"
                :long="data.ggLng"
                :heightmap="260"
                @addressData="addressDataGLoad($event)"
              />
            </div>
          </div>

        </n-card>
      </div>
      <div class="border-bottom" v-if="false">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('物业特色标签') }}
              </div>
              <n-descriptions label-placement="top" content-class="descriptions-cont" :column="1">
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">简体中文</span>
                  </template>
                  <div class="tag-div" v-if="tag_zh.length > 0">
                    <div class="tag-div-item" v-for="tags in tag_zh">
                      <n-icon size="16" color="#C4C4C4">
                        <TagOutlined/>
                      </n-icon>
                      <span style="margin-left: 5px">{{ tags }}</span>
                    </div>
                  </div>
                  <template v-else>
                    {{ translang('暂无标签') }}
                  </template>
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">日本语</span>
                  </template>
                  <div class="tag-div" v-if="tag_ja.length > 0">
                    <div class="tag-div-item" v-for="tags in tag_ja">
                      <n-icon size="16" color="#C4C4C4">
                        <TagOutlined/>
                      </n-icon>
                      <span style="margin-left: 5px">{{ tags }}</span>
                    </div>
                  </div>
                  <template v-else>
                    {{ translang('暂无标签') }}
                  </template>
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">English</span>
                  </template>
                  <div class="tag-div" v-if="tag_en.length > 0">
                    <div class="tag-div-item" v-for="tags in tag_en">
                      <n-icon size="16" color="#C4C4C4">
                        <TagOutlined/>
                      </n-icon>
                      <span style="margin-left: 5px">{{ tags }}</span>
                    </div>
                  </div>
                  <template v-else>
                    {{ translang('暂无标签') }}
                  </template>
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">한국어</span>
                  </template>
                  <div class="tag-div" v-if="tag_ko.length > 0">
                    <div class="tag-div-item" v-for="tags in tag_ko">
                      <n-icon size="16" color="#C4C4C4">
                        <TagOutlined/>
                      </n-icon>
                      <span style="margin-left: 5px">{{ tags }}</span>
                    </div>
                  </div>
                  <template v-else>
                    {{ translang('暂无标签') }}
                  </template>
                </n-descriptions-item>
                <n-descriptions-item>
                  <template #label>
                    <span class="descriptions-label">繁体中文</span>
                  </template>
                  <div class="tag-div" v-if="tag_zh_CN.length > 0">
                    <div class="tag-div-item" v-for="tags in tag_zh_CN">
                      <n-icon size="16" color="#C4C4C4">
                        <TagOutlined/>
                      </n-icon>
                      <span style="margin-left: 5px">{{ tags }}</span>
                    </div>
                  </div>
                  <template v-else>
                    {{ translang('暂无标签') }}
                  </template>
                </n-descriptions-item>
              </n-descriptions>
            </div>
            <n-button type="primary"
                      @click="editItem('TagList',translang('编辑标签'),data.tagListLanguage)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('物业描述') }}
              </div>
            </div>
            <n-button type="primary"
                      @click="editItem('Description',translang('编辑物业描述'),data.descriptionLanguage)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-descriptions label-placement="top" content-class="descriptions-cont" :column="1">
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">简体中文</span>
              </template>
              {{
                data.descriptionLanguage.zh && data.descriptionLanguage.zh.content ? data.descriptionLanguage.zh.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">日本语</span>
              </template>
              {{
                data.descriptionLanguage.ja && data.descriptionLanguage.ja.content ? data.descriptionLanguage.ja.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">English</span>
              </template>
              {{
                data.descriptionLanguage.en && data.descriptionLanguage.en.content ? data.descriptionLanguage.en.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">한국어</span>
              </template>
              {{
                data.descriptionLanguage.ko && data.descriptionLanguage.ko.content ? data.descriptionLanguage.ko.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">繁体中文</span>
              </template>
              {{
                data.descriptionLanguage.zh_CN && data.descriptionLanguage.zh_CN.content ? data.descriptionLanguage.zh_CN.content : '--'
              }}
            </n-descriptions-item>
          </n-descriptions>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('预定设置') }}
              </div>
            </div>
            <n-button type="primary" @click="handleEditYd">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-descriptions label-placement="top" content-class="descriptions-cont" :column="4">
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('货币') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.currency ? data.currency : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('语言') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{
                  data.language_name ? data.language_name : '-'
                }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('时区') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.timeZone ? data.timeZone : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('最大预定区间') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{
                  data.maxDaysNotice ? data.maxDaysNotice : '-'
                }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('最小预定区间') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{
                  data.minDaysNotice ? data.minDaysNotice : '-'
                }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('联系人') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.contactName ? data.contactName : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('联系方式') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.phone ? data.phone : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('电子邮箱') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{
                  data.contactEmail ? data.contactEmail : '-'
                }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('退房后(分钟)') }}</span>
              </template>
              <div style="margin-bottom: 10px">
                {{ data.minutesBeforeCheckin ? data.minutesBeforeCheckin : '-' }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('预定期限') }}</span>
              </template>
              <div style="margin-bottom: 10px">
                {{ data.bookingLeadTimeLabel ? data.bookingLeadTimeLabel : '-' }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('周转天数') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{
                  data.turnoverDays ? data.turnoverDays : '-'
                }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('入住前(分钟)') }}</span>
              </template>
              <div style="margin-bottom: 10px">
                {{ data.minutesAfterCheckout ? data.minutesAfterCheckout : '-' }}
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('入住时间') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.checkinAt ? data.checkinAt : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('退房时间') }}</span>
              </template>
              <div style="margin-bottom: 10px">{{ data.checkoutAt ? data.checkoutAt : '-' }}</div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">{{ translang('可见性') }}</span>
              </template>
              <div style="margin-bottom: 10px">
                <span v-if="groupIdsArr.length <= 0">全部可见</span>
                <span v-else style="color: red">部分可见</span>
              </div>
            </n-descriptions-item>
          </n-descriptions>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('物业展示图') }}
              </div>

            </div>
            <n-button type="primary" @click="imglistopen = true">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-flex v-if="imgdata.length > 0">
            <n-image
              v-for="(item, index) in imgdata"
              :src="item.fileUrl"
              :key="index"
              style="width: 60px; height: 60px; border-radius: 6px"
            />
          </n-flex>
          <div class="no-image" v-else> {{ translang('暂无任何图片') }}！</div>
        </n-card>
      </div>
      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('出行信息') }}
              </div>
            </div>
            <n-button type="primary"
                      @click="editItem('BusStation',translang('编辑出行信息'),data.busStationLanguage)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-descriptions label-placement="top" content-class="descriptions-cont" :column="1">
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">简体中文</span>
              </template>
              {{
                data.busStationLanguage.zh && data.busStationLanguage.zh.content ? data.busStationLanguage.zh.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">日本语</span>
              </template>
              {{
                data.busStationLanguage.ja && data.busStationLanguage.ja.content ? data.busStationLanguage.ja.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">English</span>
              </template>
              {{
                data.busStationLanguage.en && data.busStationLanguage.en.content ? data.busStationLanguage.en.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">한국어</span>
              </template>
              {{
                data.busStationLanguage.ko && data.busStationLanguage.ko.content ? data.busStationLanguage.ko.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">繁体中文</span>
              </template>
              {{
                data.busStationLanguage.zh_CN && data.busStationLanguage.zh_CN.content ? data.busStationLanguage.zh_CN.content : '--'
              }}
            </n-descriptions-item>
          </n-descriptions>
        </n-card>
      </div>

      <div class="border-bottom">
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('入住指南') }}
              </div>
            </div>
            <n-button type="primary"
                      @click="editItem('CheckInGuide',translang('编辑入住指南'),data.checkInGuideLanguage)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-descriptions label-placement="top" content-class="descriptions-cont" :column="1">
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">简体中文</span>
              </template>
              {{
                data.checkInGuideLanguage.zh && data.checkInGuideLanguage.zh.content ? data.checkInGuideLanguage.zh.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">日本语</span>
              </template>
              {{
                data.checkInGuideLanguage.ja && data.checkInGuideLanguage.ja.content ? data.checkInGuideLanguage.ja.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">English</span>
              </template>
              {{
                data.checkInGuideLanguage.en && data.checkInGuideLanguage.en.content ? data.checkInGuideLanguage.en.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">한국어</span>
              </template>
              {{
                data.checkInGuideLanguage.ko && data.checkInGuideLanguage.ko.content ? data.checkInGuideLanguage.ko.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">繁体中文</span>
              </template>
              {{
                data.checkInGuideLanguage.zh_CN && data.checkInGuideLanguage.zh_CN.content ? data.checkInGuideLanguage.zh_CN.content : '--'
              }}
            </n-descriptions-item>
          </n-descriptions>
        </n-card>
      </div>

      <div>
        <n-card :bordered="false">
          <div style="display:flex;justify-content: space-between">
            <div style="flex: 1;">
              <div
                style="font-size: 20px;font-weight: 500;color: rgb(31,34,37);margin-bottom: 12px">
                {{ translang('订房必读') }}
              </div>
            </div>
            <n-button type="primary"
                      @click="editItem('RequiredBook',translang('编辑订房必读'),data.requiredBookLanguage)">
              <template #icon>
                <n-icon>
                  <EditOutlined/>
                </n-icon>
              </template>
              {{ translang('编辑') }}
            </n-button>
          </div>
          <n-descriptions label-placement="top" content-class="descriptions-cont" :column="1">
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">简体中文</span>
              </template>
              {{
                data.requiredBookLanguage.zh && data.requiredBookLanguage.zh.content ? data.requiredBookLanguage.zh.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">日本语</span>
              </template>
              {{
                data.requiredBookLanguage.ja && data.requiredBookLanguage.ja.content ? data.requiredBookLanguage.ja.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">English</span>
              </template>
              {{
                data.requiredBookLanguage.en && data.requiredBookLanguage.en.content ? data.requiredBookLanguage.en.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">한국어</span>
              </template>
              {{
                data.requiredBookLanguage.ko && data.requiredBookLanguage.ko.content ? data.requiredBookLanguage.ko.content : '--'
              }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                <span class="descriptions-label">繁体中文</span>
              </template>
              {{
                data.requiredBookLanguage.zh_CN && data.requiredBookLanguage.zh_CN.content ? data.requiredBookLanguage.zh_CN.content : '--'
              }}
            </n-descriptions-item>
          </n-descriptions>
        </n-card>
      </div>
    </n-spin>

    <baseEdit ref="editBaseRef" @reloadTable="loadData"/>
    <ydEdit ref="editYdRef" @reloadTable="loadData"/>
    <ggEdit ref="editGgRef" @reloadTable="loadData"/>

    <!-- 图片详细 -->
    <a-drawer
      v-model:open="imglistopen"
      class="custom-class"
      root-class-name="root-class-name"
      style="padding: 9px"
      title="图片详细"
      width="1200"
      placement="right"
      @after-open-change="afterOpenChange"
    >
      <div class="flex-column">
        <!-- 封面图 -->
        <div class="p-5" style="border-bottom: 1px solid #ebeef5">
          <div class="flex-row">
            <img
              src="~@/assets/images/fmsyt.png"
              style="object-fit: cover; width: 274px; position: absolute; right: 8px; top: 91px"
            />
            <img
              v-if="gallery_cover"
              :src="gallery_cover"
              class="mr-3"
              style="width: 420px; height: 250px; object-fit: cover"
            />

            <div
              v-else
              class="text-c"
              style="
                width: 420px;
                height: 250px;
                background: rgba(0, 0, 0, 0.05);
                border-radius: 5px;
                margin-right: 20px;
                text-align: center;
              "
            >
              <img
                src="~@/assets/images/zwfm.png"
                style="object-fit: cover; width: 100px; margin: 82px auto"
              />
            </div>
            <div class="flex-item lh-2">
              <div class="f18 fw">封面图片</div>
              <div class="c999 f14">
                您的封面图片是客人对您的房源的第一印象。建议尺寸750*600px。
              </div>
              <!-- <div class="c999 f14"> 从下方图片列表中选择一张成为封面图 </div> -->
            </div>
          </div>
        </div>

        <!-- 所有图片 -->
        <div class="p-5 lh-2 flex-item" style="position: relative">
          <div class="f18 fw">所有图片</div>
          <div class="c999 f12">拖放图片能更改图片显示的顺序</div>
          <div style="position: absolute; top: 10px; right: 10px">
            <FileChooser
              :maxNumber="10"
              v-model:value="data.gallery_images"
              :nolist="true"
              @change="FileChooserchange"
            />
          </div>

          <div class="mt-5">
            <Draggable
              animation="300"
              :list="imgdata"
              group="people"
              itemKey="fileUrl"
              @update="dragchange"
            >
              <template #item="{ element }">
                <div class="imgitem mr-3 lh-3 mb-3">
                  <n-image-group>
                    <n-space>
                      <n-image object-fit="cover" :width="260" :src="element.fileUrl"/>
                    </n-space>
                  </n-image-group>
                  <div class="c333 flex-row f12">
                    <div class="nohuanahang flex-item mr-2" style="overflow-x: hidden">{{
                        element.name
                      }}
                    </div>
                    <div class="text-r c999">
                      <a class="mr-2 c999" @click="delimg(element.fileUrl)"> 删除 </a>
                      <a class="mr-2 c999" @click="coverimg(element.fileUrl)">
                        <n-icon size="14" color="#999">
                          <Reload/>
                        </n-icon>
                        设为封面</a>
                    </div>
                  </div>
                </div>
              </template>
            </Draggable>
          </div>
        </div>
      </div>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
import {EditOutlined, TagOutlined} from '@vicons/antd';
import {Reload} from '@vicons/ionicons5';
import {onMounted, reactive, ref, watch} from "vue";
import baiduMap from '@/views/smjcomm/BMap.vue';
import googleMap from '@/views/smjcomm/GMap.vue';
import BMapView from '@/views/smjcomm/BMapView.vue';
import GMapView from '@/views/smjcomm/GMapView.vue';
import {jsontoobj, translang} from "@/utils/smjcomm";
import {Edit, getGallery, setGalleryCover, View} from "@/api/pmsProperty";
import {useRouter} from "vue-router";
import {useGlobSetting} from "@/hooks/setting";
import {useUserStoreWidthOut} from "@/store/modules/user";
import {useMessage} from "naive-ui";
import {ResultEnum} from '@/enums/httpEnum';
import componentSetting from '@/settings/componentSetting';
import {Attachment} from '@/components/FileChooser/src/model';
import baseEdit from '@/views/pmsProperty/comm/edit_base.vue';
import ydEdit from '@/views/pmsProperty/comm/edit_yd.vue';
import ggEdit from '@/views/pmsProperty/comm/edit_gg.vue';
import Draggable from "vuedraggable";
import FileChooser from "@/components/FileChooser/index.vue";

const show = ref(false);
const data = ref({
  bookingClose: false,
  leaseClose: false,
  batchReservation: false,
  groupIds: '',
  cover: '',
  nameLanguage: {},
  addressLanguage: {},
  descriptionLanguage: {},
  lat: '',
  lng: '',
  addressDetail: '',
  ggLat: '',
  ggLng: '',
  ggAddressDetail: '',
  currency: '',
  language_name: '',
  timeZone: '',
  maxDaysNotice: '',
  minDaysNotice: '',
  contactName: '',
  phone: '',
  contactEmail: '',
  minutesBeforeCheckin: '',
  bookingLeadTimeLabel: '',
  turnoverDays: '',
  minutesAfterCheckout: '',
  checkinAt: '',
  checkoutAt: '',
  gallery_images: [],
  tagListLanguage: {},
  busStationLanguage: {},
  checkInGuideLanguage: {},
  requiredBookLanguage: {},
  regionNameLanguage:{},
  regionId: 0,
});
let tag_zh = ref([]);
let tag_en = ref([]);
let tag_ko = ref([]);
let tag_ja = ref([]);
let tag_zh_CN = ref([]);
let imgdata = ref([]);
let groupIdsArr = ref([]);
let gallery_cover = ref('');
let editMap = ref(false);
let bookingClose = ref(false);
let leaseClose = ref(false);
let batchReservation = ref(false);
let mapview = ref('google');
let mapkey = ref(1);
let imglistopen = ref(false);
const router = useRouter();
const globSetting = useGlobSetting();
const {uploadUrl} = globSetting;
const urlPrefix = globSetting.urlPrefix || '';
const useUserStore = useUserStoreWidthOut();
const uploadHeaders = reactive({
  Authorization: useUserStore.token,
  uploadType: 'image',
});
const message = useMessage();
const editBaseRef = ref();
const editYdRef = ref();
const editGgRef = ref();
const bookingCloseloadingRef = ref(false)
const leaseCloseloadingRef = ref(false)
const batchReservationloadingRef = ref(false)

onMounted(() => {
  loadData();
  getGalleryLoad();
});


function loadData() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    View({id: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id})
      .then((res) => {
        data.value = res
        groupIdsArr.value = data.value.groupIds ? data.value.groupIds.split(",").map(item => Number(item)) : [];
        if (res.nameLanguage) {
          data.value.nameLanguage = jsontoobj(res.nameLanguage);
          if (data.value.nameLanguage.zh_CN) {
            data.value.nameLanguage.zh_CN = data.value.nameLanguage.zh_CN
          } else if (data.value.nameLanguage.zh_cn) {
            data.value.nameLanguage.zh_CN = data.value.nameLanguage.zh_cn
          } else {
            data.value.nameLanguage.zh_CN = {}
          }
        } else {
          data.value.nameLanguage = {};
        }

        if (res.regionInfo && res.regionInfo.regionNameLanguage){
          data.value.regionNameLanguage = jsontoobj(res.regionInfo.regionNameLanguage);
          if (data.value.regionNameLanguage.zh_CN) {
            data.value.regionNameLanguage.zh_CN = data.value.regionNameLanguage.zh_CN
          } else if (data.value.regionNameLanguage.zh_cn) {
            data.value.regionNameLanguage.zh_CN = data.value.regionNameLanguage.zh_cn
          } else {
            data.value.regionNameLanguage.zh_CN = {}
          }
        } else {
          data.value.regionNameLanguage = {};
        }
console.log('data11', data)
        if (res.addressLanguage) {
          data.value.addressLanguage = jsontoobj(res.addressLanguage);
          if (data.value.addressLanguage.zh_CN) {
            data.value.addressLanguage.zh_CN = data.value.addressLanguage.zh_CN
          } else if (data.value.addressLanguage.zh_cn) {
            data.value.addressLanguage.zh_CN = data.value.addressLanguage.zh_cn
          } else {
            data.value.addressLanguage.zh_CN = {}
          }
        } else {
          data.value.addressLanguage = {};
        }

        if (res.tagListLanguage) {
          let tagListLanguageObj = jsontoobj(res.tagListLanguage); //多语言物业标签
          data.value.tagListLanguage = tagListLanguageObj
          tag_zh.value = tagListLanguageObj.zh.content ? JSON.parse(tagListLanguageObj.zh.content) : []
          tag_en.value = tagListLanguageObj.en.content ? JSON.parse(tagListLanguageObj.en.content) : []
          tag_ko.value = tagListLanguageObj.ko.content ? JSON.parse(tagListLanguageObj.ko.content) : []
          tag_ja.value = tagListLanguageObj.ja.content ? JSON.parse(tagListLanguageObj.ja.content) : []
          if (tagListLanguageObj.zh_cn && tagListLanguageObj.zh_cn.content) {
            tag_zh_CN.value = JSON.parse(tagListLanguageObj.zh_cn.content)
          } else if (tagListLanguageObj.zh_CN && tagListLanguageObj.zh_CN.content) {
            tag_zh_CN.value = JSON.parse(tagListLanguageObj.zh_CN.content)
          } else {
            tag_zh_CN.value = [];
          }
        } else {
          data.value.tagListLanguage = {}
          tag_zh.value = []
          tag_en.value = []
          tag_ko.value = []
          tag_ja.value = []
          tag_zh_CN.value = []
        }
        if (res.language) {
          if (res.language == 'zh') {
            data.value.language_name = '简体中文'
          } else if (res.language == 'en') {
            data.value.language_name = 'English'
          } else if (res.language == 'ko') {
            data.value.language_name = '한국어'
          } else if (res.language == 'ja') {
            data.value.language_name = '日本语'
          } else if (res.language == 'zh_CN') {
            data.value.language_name = '繁体中文'
          }
        }
        if (res.descriptionLanguage) {
          data.value.descriptionLanguage = jsontoobj(res.descriptionLanguage);
          if (data.value.descriptionLanguage.zh_CN) {
            data.value.descriptionLanguage.zh_CN = data.value.descriptionLanguage.zh_CN
          } else if (data.value.descriptionLanguage.zh_cn) {
            data.value.descriptionLanguage.zh_CN = data.value.descriptionLanguage.zh_cn
          } else {
            data.value.descriptionLanguage.zh_CN = {}
          }
        } else {
          data.value.descriptionLanguage = {};
        }

        if (res.busStationLanguage) {
          data.value.busStationLanguage = jsontoobj(res.busStationLanguage);
          if (data.value.busStationLanguage.zh_CN) {
            data.value.busStationLanguage.zh_CN = data.value.busStationLanguage.zh_CN
          } else if (data.value.busStationLanguage.zh_cn) {
            data.value.busStationLanguage.zh_CN = data.value.busStationLanguage.zh_cn
          } else {
            data.value.busStationLanguage.zh_CN = {}
          }
        } else {
          data.value.busStationLanguage = {};
        }

        if (res.checkInGuideLanguage) {
          data.value.checkInGuideLanguage = jsontoobj(res.checkInGuideLanguage);
          if (data.value.checkInGuideLanguage.zh_CN) {
            data.value.checkInGuideLanguage.zh_CN = data.value.checkInGuideLanguage.zh_CN
          } else if (data.value.checkInGuideLanguage.zh_cn) {
            data.value.checkInGuideLanguage.zh_CN = data.value.checkInGuideLanguage.zh_cn
          } else {
            data.value.checkInGuideLanguage.zh_CN = {}
          }
        } else {
          data.value.checkInGuideLanguage = {};
        }

        if (res.requiredBookLanguage) {
          data.value.requiredBookLanguage = jsontoobj(res.requiredBookLanguage);
          if (data.value.requiredBookLanguage.zh_CN) {
            data.value.requiredBookLanguage.zh_CN = data.value.requiredBookLanguage.zh_CN
          } else if (data.value.requiredBookLanguage.zh_cn) {
            data.value.requiredBookLanguage.zh_CN = data.value.requiredBookLanguage.zh_cn
          } else {
            data.value.requiredBookLanguage.zh_CN = {}
          }
        } else {
          data.value.requiredBookLanguage = {};
        }
        data.value.leaseClose = res.leaseClose == 1
        data.value.bookingClose = res.bookingClose == 1
        data.value.batchReservation = res.batchReservation == "Y"

      })
      .finally(() => {
        show.value = false;
      });
  });
}

const getGalleryLoad = () => {
  getGallery({uid: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id}).then((res) => {
    imgdata.value = res.gallery_images ? JSON.parse(res.gallery_images) : [];
    gallery_cover.value = res.gallery_cover ? res.gallery_cover : '';
    data.value.gallery_images = res.gallery_images
      ? JSON.parse(res.gallery_images).map((m) => {
        return m.fileUrl;
      })
      : [];
  });
};

watch(
  () => useUserStore.getuserPms,
  (_newVal, _oldVal) => {
    if (!router.currentRoute.value.query.id) {
      loadData();
      getGalleryLoad();
    }
  },
  {
    immediate: true,
    deep: true,
  }
);

//上传结束
function finishUploadCover({event: Event}) {
  const res = JSON.parse(Event.target.response);
  const infoField = componentSetting.upload.apiSetting.infoField;
  const {code} = res;
  const msg = res.msg || res.message || '上传失败';
  const result = res[infoField] as Attachment;

  //成功
  if (code === ResultEnum.SUCCESS) {
    Edit({
      id: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id,
      cover: result.fileUrl
    }).then((_res) => {
      message.success('上传' + result.name + '成功');
      data.value.cover = result.fileUrl
    });

  } else {
    message.error(msg);
  }
}

const addressDataLoad = (res) => {
  console.log('addressData', res);
  data.value.lat = res.lat + '';
  data.value.lng = res.lng + '';
  data.value.addressDetail = res.myValuetext;
  Edit({
    id: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id,
    lat: res.lat,
    lng: res.lng,
    addressDetail: res.myValuetext,
  }).then((_res) => {
    message.success('操作成功');
    editMap.value = false;
  });
};
const addressDataGLoad = (res) => {
  console.log('addressData', res);
  data.value.ggLat = res.lat + '';
  data.value.ggLng = res.lng + '';
  data.value.ggAddressDetail = res.myValuetext;
  Edit({
    id: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id,
    ggLat: res.lat,
    ggLng: res.lng,
    ggAddressDetail: res.myValuetext,
  }).then((_res) => {
    message.success('操作成功');
    editMap.value = false;
  });
};

const afterOpenChange = (bool: boolean) => {
  console.log('open', bool);
};

const FileChooserchange = (res) => {
  imgdata.value = imgdata.value.concat(res.allvalue);

  setGallery();
};

const dragchange = () => {
  data.value.gallery_images = imgdata.value.map((m) => {
    return m.fileUrl;
  });
  setGallery();
};

const delimg = (url) => {
  imgdata.value = imgdata.value
    .filter((f) => {
      if (f.fileUrl != url) {
        return f;
      }
    })
    .map((m) => {
      return m;
    });
  setGalleryCover({
    uid: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id,
    gallery_cover: gallery_cover.value,
    gallery_images: JSON.stringify(imgdata.value),
  }).then((res) => {
    getGalleryLoad();
  });
};

const coverimg = (url) => {
  gallery_cover.value = url;
  setGallery();
};

const setGallery = () => {
  setGalleryCover({
    uid: router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id,
    gallery_cover: gallery_cover.value,
    gallery_images: JSON.stringify(imgdata.value),
  }).then((res) => {
    getGalleryLoad();
  });
};

function editItem(type, title, data, regionId = 0) {
  // console.log('editItem', type, title, data, regionId);
  editBaseRef.value.openModal(router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id, type, title, data, regionId);
}

function handleEditYd() {
  editYdRef.value.openModal(data.value);
}

function handleEditGg() {
  editGgRef.value.openModal(data.value);
}

function handleChangeleaseClose (params){
  leaseCloseloadingRef.value = true
  Edit({"leaseClose":params ? 1 : 2,"id":router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id})
    .then((_res) => {
      leaseCloseloadingRef.value = false
      message.success('操作成功');
    })
    .catch((err) => {
      leaseCloseloadingRef.value = false
      message.error(err);
    });
}

function handleChangebookingClose (params){
  bookingCloseloadingRef.value = true
  Edit({"bookingClose":params ? 1 : 2,"id":router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id})
    .then((_res) => {
      bookingCloseloadingRef.value = false
      message.success('操作成功');
    })
    .catch((err) => {
      bookingCloseloadingRef.value = false
      message.error(err);
    });
}

function handleChangebatchReservation (params){
  data.value.batchReservation = params
  batchReservationloadingRef.value = true
  Edit({"batchReservation":params ? "Y" : "N","id":router.currentRoute.value.query.id ? router.currentRoute.value.query.id : useUserStore.getuserPms.id})
    .then((_res) => {
      batchReservationloadingRef.value = false
      message.success('操作成功');
    })
    .catch((err) => {
      batchReservationloadingRef.value = false
      message.error(err);
    });
}
</script>

<style lang="less">
.img-cover {
  position: relative;
  width: 133px;
  height: 200px;

  img {
    display: block;
    width: 100%;
    height: 100%;
  }

  a {
    display: flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(5, 61, 200, 0.67);
    height: 32px;

    span {
      font-size: 14px;
      color: #fff;
      margin-left: 3px;
    }
  }
}

.border-bottom {
  border-bottom: 1px solid #EEEEEE;
}

.descriptions-label {
  color: #707070;
  font-size: 12px;
  font-weight: normal;
}

.descriptions-cont {
  font-size: 14px;
  font-weight: bold;
}

.map-location-info {
  font-size: 14px;
  color: #707070;
  font-weight: bold;
  margin-bottom: 10px;

  span {
    margin-right: 20px;

    span {
      color: #3D3D3D;
    }
  }
}

.mapdiv {
  position: absolute;
  top: 15px;
  left: 15px;
  height: 38px;
  text-align: center;
  color: #3d3d3d;
  padding: 8px 15px;
  background-color: #fff;
  box-shadow: 4px 3px 5px #999999a8;
  z-index: 300;

  .tabview {
    color: #333;
    font-size: 14px;

    .showmap {
      color: #053dc8;
      font-weight: 550;
    }

    img {
      width: 18px;
      height: 18px;
      float: left;
      margin-right: 6px;
    }
  }
}

.mapbtm {
  position: absolute;
  z-index: 300;
  bottom: 0;
  text-align: left;
  width: 100%;
  background: #6a7b8c38;
  padding: 5px 10px;
  font-weight: 550;
}

.tag-div {
  display: flex;
  align-items: center;
}

.tag-div-item {
  display: flex;
  align-items: center;
  margin-right: 20px;
}

.no-image {
  font-size: 14px;
}

.imgitem {
  display: inline-block;
  width: 260px;

  img {
    height: 150px;
    border-radius: 5px;
  }
}
</style>
