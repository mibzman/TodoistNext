# Todoist Next
> Enable a global overview of the 'next' task in all projects.

This project is very customized for my personal use, and is generally not ready for others to use without modifications.

The principal is you provide some 'root' projects, and TodoistNext looks at all subchildren of those projects.  It then retrieves the task nearest the top of the project that does not have a specefic assigned data. Then it applies the 'next' tag, which can be viewed with a filter.

On my homeserver, I have a chron job that runs this script at the top of every hour, so it's reasonably up to date.

Additionally this project also contains some sync features with notion.  At the time, I maintained a link of interesting links for future use.  I wanted to keep them in notion, but adding them to notion was cumbersome.  So I create a 'notion' todoist project, and this script would move the contents into a notion table.  I think the only reason I did this was because I already had the todoist injest working for the next feature.

Also I promise I use git for everything, but I had to scrub the history since it used to contain secrets.
