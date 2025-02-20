# grokengage

Social Task Reward Program: "Engage & Earn" (Enhanced) Built 100% with GrokAi
Program Overview
"Engage & Earn" is a points-based system where users complete social media tasks and a daily check-in to earn points. These points can be redeemed for rewards like discounts, exclusive content, badges, or prizes. The addition of a daily check-in task boosts retention and habitual use.



engage-and-earn/
├── main.go           # Entry point and server setup
├── models/           # Database models
│   ├── user.go       # User model
│   ├── task.go       # Task model
│   ├── checkin.go    # Check-in model
│   ├── reward.go     # Reward model
├── handlers/         # API handlers
│   ├── user.go       # User-related endpoints
│   ├── task.go       # Task-related endpoints
│   ├── checkin.go    # Check-in endpoints
│   ├── reward.go     # Reward-related endpoints
├── db/               # Database connection
│   └── db.go         # PostgreSQL setup with GORM
└── go.mod            # Module file


Prerequisites
Install Go (1.21+ recommended)
Install PostgreSQL and create a database named engage_and_earn
Run go get to fetch dependencies after setting up go.mod




Core Components
Tasks
Users complete social media actions and a daily check-in, each with assigned point values based on effort and value:
Daily Check-In: Log into the program dashboard once per day (e.g., click a “Check-In” button).
Points: 2
(Incentive: Encourages daily visits; low effort, consistent reward.)
Follow a Profile: Follow a specified account (e.g., brand, influencer).
Points: 5
Like a Post: Like a designated post (provided via link).
Points: 3
Repost/Share a Post: Share or repost specific content to their feed.
Points: 10
Visit a Page: Click a unique link to visit a webpage (tracked via analytics).
Points: 2
Bonus Task: Comment on a Post: Leave a meaningful comment (min. 5 words) on a specified post.
Points: 8
Points System
Points awarded instantly upon task completion, verified via APIs or user-submitted proof.
Daily Cap: Max 50 points/day from tasks (excludes streak bonuses—see below).
Check-In Streak Bonus:  
5 consecutive days: +5 bonus points (total 15 points over 5 days).  
10 consecutive days: +15 bonus points (total 40 points over 10 days).  
Resets if a day is missed, encouraging consistency.
Task Completion Bonus: Complete all daily tasks (excluding check-in) for an extra 10 points.
Reward Tiers
Updated to reflect daily check-in accumulation:
25 Points: Virtual badge (e.g., “Daily Dynamo”) for profile display.  
75 Points: Exclusive digital content (e.g., wallpaper, short video).  
200 Points: 10% discount code for a partner brand.  
400 Points: Raffle entry for a physical prize (e.g., $50 gift card).  
800 Points: Premium reward (e.g., personalized shoutout or merchandise like a mug).
User Interface
Dashboard Features:  
Daily Check-In button (turns green once completed).  
Streak tracker (e.g., “Day 3 of 5”) with progress bar.  
Available tasks with point values and deadlines.  
Points balance and reward redemption catalog.
Tasks refresh daily or weekly; check-in is a constant feature.
Verification Process
Daily Check-In: Auto-verified upon dashboard login and button click (tied to user account).  
Social Tasks:  
Automated via APIs (e.g., X API for follows/likes/reposts).  
Manual for comments/page visits (users submit screenshots or use trackable links).
Anti-Cheat: One check-in per 24-hour cycle, IP/account monitoring, minimum account age (30 days).



How It Works
Sign-Up: Users register with their social media handle(s) and create an account.  
Daily Routine:  
Log in, click “Check-In” for 2 points.  
Browse and complete tasks (e.g., “Repost 
@xAI
 update for 10 points”).
Verification: Check-in is instant; tasks verified automatically or via proof.  
Points Awarded: Credited to account with streak bonuses applied at milestones.  
Redemption: Users browse the catalog and redeem points for rewards.




Example Workflow
Day 1 Tasks:  
Check-In (2 points)  
Follow 
@BrandX
 (5 points)  
Like post [link] (3 points)  
Repost update [link] (10 points)  
Visit www.brandx.com/promo (2 points)
User Action: Completes all tasks.  
Result: 22 points + 10-point task completion bonus = 32 points.  
Day 5: After 5 daily check-ins (10 points) + streak bonus (5 points) + tasks, user might have ~150 points.  
Progress: Redeems 75 points for digital content, keeps earning.


Benefits
For Users: Daily check-ins add a low-effort way to earn points, while streaks and bonuses gamify participation.  
For Platform: Higher retention (daily visits), increased social engagement, and more consistent traffic.  
Scalability: Check-in feature integrates seamlessly with existing tasks.
Enhanced Features
Streak Milestones: Visual rewards (e.g., badge upgrades at 10, 30, 60 days).  
Random Daily Bonus: 1 in 10 chance for check-in to award 5 points instead of 2.  
Leaderboard: Top check-in streaks or point earners monthly get extra perks.
