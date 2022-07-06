# Knowledge Forum Plugin for Steampipe

Use SQL to query Knowledge Forum database

## Usage

Prerequisites:

- git client
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/macc704/steampipe-plugin-kf.git
cd steampipe-plugin-kf
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

- please check your target communityId from [CommunityId Finder](https://macc704.github.io/kf6apijs/community_finder/)

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/kf.spc
```

kf.spc file
```code
connection "kf" {
  plugin = "kf"

  username = "input your kf username"
  password = "input your kf password"
  url = "input your kf site like https://kf6-stage.ikit.org/"
  communityId = "input your kf community Id"
}
```

Try it!

```shell
steampipe query
> .inspect kf_notes
> .inspect kf_views
> .inspect kf_links
> .inspect kf_authors
```

Run a query:

```sql
select kf_authors.last_name, count(*) from kf_notes inner join kf_authors on kf_notes.author = kf_authors.id group by kf_authors.last_name order by count(*) desc
```

```sql
select kf_views.title, count(*) from kf_views inner join kf_links on kf_links.from = kf_views.id inner join kf_notes on kf_links.to = kf_notes.id group by kf_views.title order by count(*) desc
```

```sql
select kf_authors.name, kf_views.title, kf_notes.title from kf_views inner join kf_links on kf_links.from = kf_views.id inner join kf_notes on kf_links.to = kf_notes.id inner join kf_authors on kf_notes.author = kf_authors.id where kf_views.title like '%xxx%'
```