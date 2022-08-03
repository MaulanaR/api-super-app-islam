package translations

import "grest.dev/grest"

func AddEnUS() {
	grest.AddTranslation("en-US", map[string]string{
		"400_bad_request":                     "The request cannot be performed because of malformed or missing parameters.",
		"401_unauthorized":                    "Unauthorized. Please Re-Login",
		"403_forbidden":                       "The user does not have permission to :action.",
		"404_not_found":                       "The resource you have specified cannot be found.",
		"500_internal_error":                  "Failed to connect to the server, please try again later.",
		"deleted":                             "Data has been deleted.",
		"not_found":                           ":entity data with :key = :value cannot be found.",
		"greater_than":                        "The value of a :key must be greater than :value",
		"less_than":                           "The value of a :key must be less than :value",
		"not_in":                              "The file extension must be one of the following types: (:value).",
		"storage_limit":                       "Cannot upload file because of your storage limit. Please upgrade limit storage.",
		"required_value":                      "The :key field is required.",
		"unique":                              "The :attribute (:value) has already been taken.",
		"success":                             "Success.",
		"attendance_is_exists":                "Attendance for selected date is already exist.",
		"invalid_operating_hour":              "Invalid Operating Hours",
		"invalid_operating_day":               "Invalid Operating Days",
		"invalid_username_or_password":        "Invalid username or password",
		"insufficient_timeoff_balance":        "Insufficient time off's balance",
		"id":                                  "id",
		"code":                                "code",
		"branches":                            "branch data",
		"employees":                           "employee data",
		"data_stores.departments.create":      "create department data",
		"timeschedules_in_used_on_attendance": "This time schedule has been used in attendance",
		"invalid_data_type":                   "Invalid data type or length. Please check your parameters.",
	})
}
