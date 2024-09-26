# Get the current user Path
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

# Define the new path you want to add
$newPath = "C:\Program Files (x86)\GnuWin32\bin"

# Check if the new path already exists in the Path
if (-not ($userPath -split ';' | ForEach-Object { $_.Trim() } | Where-Object { $_ -eq $newPath })) {
    # Concatenate the new path to the existing user path
    $updatedPath = "$userPath;$newPath"

    # Set the updated path as the user Path variable
    [Environment]::SetEnvironmentVariable("Path", $updatedPath, "User")

    Write-Output "Path updated successfully."
} else {
    Write-Output "The path '$newPath' is already in the user Path."
}