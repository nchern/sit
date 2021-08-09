#!/bin/sh -ue


printf "# Open issues\n\n\`\`\`\n"
sit  list | awk '/OPEN/ {$2="";$3=""; $4=""; print }'
printf "\`\`\`\n"
