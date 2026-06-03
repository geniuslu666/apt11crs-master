package airhousePublicApi

import "time"

type Stay struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type RoomType struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type RoomUnit struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type RatePlan struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type ConnectedRatePlan struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type Folio struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type Paginator struct {
	TotalCount  int    `json:"total_count"`
	TotalPages  int    `json:"total_pages"`
	CurrentPage int    `json:"current_page"`
	PageSize    int    `json:"page_size"`
	PrevPageURL string `json:"prev_page_url"`
	NextPageURL string `json:"next_page_url"`
}
type Property struct {
	ID  interface{} `json:"id,omitempty"`
	UID interface{} `json:"uid,omitempty"`
}
type RoomReservation struct {
	UID interface{} `json:"uid,omitempty"`
}
type RoomReservations struct {
	ID                 interface{}       `json:"id,omitempty"`
	ObjectType         string            `json:"object_type"`
	UID                interface{}       `json:"uid,omitempty"`
	Stay               Stay              `json:"stay"`
	Property           Property          `json:"property"`
	RoomType           RoomType          `json:"room_type"`
	RoomUnit           RoomUnit          `json:"room_unit"`
	RatePlan           RatePlan          `json:"rate_plan"`
	ConnectedRatePlan  ConnectedRatePlan `json:"connected_rate_plan"`
	Folio              Folio             `json:"folio"`
	MainGuest          MainGuest         `json:"main_guest"`
	CheckinDate        string            `json:"checkin_date"`
	CheckoutDate       string            `json:"checkout_date"`
	CheckinTime        string            `json:"checkin_time"`
	CheckoutTime       string            `json:"checkout_time"`
	Status             string            `json:"status"`
	CheckinStatus      string            `json:"checkin_status"`
	AdultCount         int               `json:"adult_count"`
	ChildCount         int               `json:"child_count"`
	InfantCount        int               `json:"infant_count"`
	GuestRemarks       string            `json:"guest_remarks"`
	BookingFee         float64           `json:"booking_fee"`
	ChannelFee         float64           `json:"channel_fee"`
	CleaningFee        float64           `json:"cleaning_fee"`
	CancellationFee    float64           `json:"cancellation_fee"`
	TotalChargeAmount  float64           `json:"total_charge_amount"`
	TotalPaymentAmount float64           `json:"total_payment_amount"`
	OutstandingBalance float64           `json:"outstanding_balance"`
	PrepaidAmount      float64           `json:"prepaid_amount"`
	Charges            []Charges         `json:"charges"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}
type Charges struct {
	ID          interface{} `json:"id,omitempty"`
	UID         interface{} `json:"uid,omitempty"`
	Date        string      `json:"date"`
	FeeType     string      `json:"fee_type"`
	Amount      float64     `json:"amount"`
	Description interface{} `json:"description"`
}
type Booker struct {
	ID            interface{} `json:"id,omitempty"`
	ObjectType    interface{} `json:"object_type"`
	UID           interface{} `json:"uid,omitempty"`
	FirstName     interface{} `json:"first_name"`
	LastName      interface{} `json:"last_name"`
	FirstNameKana interface{} `json:"first_name_kana"`
	LastNameKana  interface{} `json:"last_name_kana"`
	FullName      interface{} `json:"full_name"`
	Language      interface{} `json:"language"`
	Email         interface{} `json:"email"`
	Phone         interface{} `json:"phone"`
	Nationality   interface{} `json:"nationality"`
	Address       interface{} `json:"address"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
type MainGuest struct {
	ID              interface{} `json:"id,omitempty"`
	ObjectType      interface{} `json:"object_type"`
	UID             interface{} `json:"uid,omitempty"`
	FirstName       interface{} `json:"first_name"`
	LastName        interface{} `json:"last_name"`
	FirstNameKana   interface{} `json:"first_name_kana"`
	LastNameKana    interface{} `json:"last_name_kana"`
	FullName        interface{} `json:"full_name"`
	Language        interface{} `json:"language"`
	Email           interface{} `json:"email"`
	Phone           interface{} `json:"phone"`
	Nationality     interface{} `json:"nationality"`
	Address         interface{} `json:"address"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	RoomReservation struct {
		ID string `json:"id"`
	} `json:"room_reservation,omitempty"`
}
type StayData struct {
	ID                    string             `json:"id,omitempty"`
	ObjectType            string             `json:"object_type"`
	UID                   string             `json:"uid,omitempty"`
	Property              Property           `json:"property"`
	ReservationNumber     string             `json:"reservation_number"`
	ReservationSourceCode string             `json:"reservation_source_code"`
	ReservationSourceName string             `json:"reservation_source_name"`
	Booker                Booker             `json:"booker"`
	RoomReservations      []RoomReservations `json:"room_reservations"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
}
type RoomUnits struct {
	ID         interface{} `json:"id,omitempty"`
	ObjectType string      `json:"object_type"`
	UID        interface{} `json:"uid,omitempty"`
	RoomNo     string      `json:"room_no"`
}
type RoomTypes struct {
	ID         interface{} `json:"id,omitempty"`
	ObjectType string      `json:"object_type"`
	UID        interface{} `json:"uid,omitempty"`
	Name       string      `json:"name"`
	RoomUnits  []RoomUnits `json:"room_units"`
}
type RatePlans struct {
	ID                       interface{}          `json:"id,omitempty"`
	ObjectType               string               `json:"object_type"`
	UID                      interface{}          `json:"uid,omitempty"`
	Name                     string               `json:"name"`
	Connected                bool                 `json:"connected"`
	ConnectedRatePlan        interface{}          `json:"connected_rate_plan"`
	ConnectedRatePlanName    interface{}          `json:"connected_rate_plan_name"`
	PricingModel             string               `json:"pricing_model"`
	PerDayPricing            PerDayPricing        `json:"per_day_pricing"`
	ServiceFeesPerStay       []ServiceFeesPerStay `json:"service_fees_per_stay"`
	DefaultMinLos            int                  `json:"default_min_los"`
	DefaultMaxLos            int                  `json:"default_max_los"`
	MinAdvancedBookingOffset interface{}          `json:"min_advanced_booking_offset"`
	MaxAdvancedBookingOffset interface{}          `json:"max_advanced_booking_offset"`
	RatePlanLinkage          interface{}          `json:"rate_plan_linkage"`
	UpdatedAt                time.Time            `json:"updated_at"`
}
