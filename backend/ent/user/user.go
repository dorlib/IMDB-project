// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldBirthDay holds the string denoting the birthday field in the database.
	FieldBirthDay = "birth_day"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldSignupAt holds the string denoting the signup_at field in the database.
	FieldSignupAt = "signup_at"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeMovies holds the string denoting the movies edge name in mutations.
	EdgeMovies = "movies"
	// EdgeDirectors holds the string denoting the directors edge name in mutations.
	EdgeDirectors = "directors"
	// EdgeAchievements holds the string denoting the achievements edge name in mutations.
	EdgeAchievements = "achievements"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "user_reviews"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "user_comments"
	// LikesTable is the table that holds the likes relation/edge. The primary key declared below.
	LikesTable = "user_likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// MoviesTable is the table that holds the movies relation/edge.
	MoviesTable = "movies"
	// MoviesInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	MoviesInverseTable = "movies"
	// MoviesColumn is the table column denoting the movies relation/edge.
	MoviesColumn = "user_id"
	// DirectorsTable is the table that holds the directors relation/edge.
	DirectorsTable = "directors"
	// DirectorsInverseTable is the table name for the Director entity.
	// It exists in this package in order to avoid circular dependency with the "director" package.
	DirectorsInverseTable = "directors"
	// DirectorsColumn is the table column denoting the directors relation/edge.
	DirectorsColumn = "user_id"
	// AchievementsTable is the table that holds the achievements relation/edge.
	AchievementsTable = "achievements"
	// AchievementsInverseTable is the table name for the Achievement entity.
	// It exists in this package in order to avoid circular dependency with the "achievement" package.
	AchievementsInverseTable = "achievements"
	// AchievementsColumn is the table column denoting the achievements relation/edge.
	AchievementsColumn = "user_achievements"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldNickname,
	FieldDescription,
	FieldPassword,
	FieldEmail,
	FieldBirthDay,
	FieldProfile,
	FieldCountry,
	FieldGender,
	FieldSignupAt,
}

var (
	// LikesPrimaryKey and LikesColumn2 are the table columns denoting the
	// primary key for the likes relation (M2M).
	LikesPrimaryKey = []string{"user_id", "like_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
