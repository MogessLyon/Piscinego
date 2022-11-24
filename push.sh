echo "Enter the filename"
read filename
git add $filename
git commit -m "$filename"
git push
