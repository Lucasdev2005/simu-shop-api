# handler
echo "--------------------------------------------------------------------------------------"
echo "Running unit tests from handler"
echo "--------------------------------------------------------------------------------------"
go test ./internal/handler

# Repository
echo "--------------------------------------------------------------------------------------"
echo "Running unit tests from repository"
echo "--------------------------------------------------------------------------------------"

rm internal/repository/gorm.db
go test ./internal/repository