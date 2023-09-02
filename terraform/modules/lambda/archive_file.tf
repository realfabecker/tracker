data "archive_file" "zip" {
  source_file = var.zip_source_file
  type        = "zip"
  output_path = "${var.zip_output_path}/${var.zip_file_name}"
}
