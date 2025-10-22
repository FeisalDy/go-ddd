| From → To                    | Allowed?       | Reason                              |
| ----------------------------- | -------------- | ----------------------------------- |
| Interface → Application      | ✅             | Interface calls use cases           |
| Interface → Domain           | ✅ (sometimes) | For simple queries or read models   |
| Application → Domain         | ✅             | Normal use case logic               |
| Infrastructure → Application | ✅             | Implements abstractions from app    |
| Infrastructure → Domain      | ✅             | Implements repository interfaces    |
| Application → Infrastructure | ❌             | Breaks inversion                    |
| Domain → Application         | ❌             | Core should not depend on use cases |
| Domain → Infrastructure      | ❌             | Never tie business logic to DB/tech |
